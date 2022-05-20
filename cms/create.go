package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-github/v44/github"
)

func createPost() {
	token, err := login()

	if err != nil || len(token) == 0 {
		log.Fatalf("login failed: %v", err)
	}

	cli := github.NewClient(nil)

	issues, _, err := cli.Issues.
		ListByRepo(context.Background(), owner, repo, &github.IssueListByRepoOptions{
			State:  "open",
			Labels: []string{label},
			Sort:   "created",
			ListOptions: github.ListOptions{
				Page:    1,
				PerPage: 1,
			},
		})
	if err != nil {
		log.Fatalf("list issues failed: %v", err)
	}

	if len(issues) == 0 {
		log.Fatalln("no issue found")
	}

	req := resty.New()

	issue := issues[0]
	labelUrl := issue.GetLabelsURL()

	resp, err := req.R().Get(labelUrl[:len(labelUrl)-7])
	if err != nil {
		log.Fatalf("get labels failed: %v", err)
	}

	var labels []Label
	if err := json.Unmarshal(resp.Body(), &labels); err != nil || len(labels) == 0 {
		log.Fatalf("unmarshal labels failed: %v", err)
	}

	reg := regexp.MustCompile("\\[([a-zA-Z0-9-_!]+)\\] ([\u4e00-\u9fa5_a-zA-Z0-9!\\s]+)")
	params := reg.FindStringSubmatch(issue.GetTitle())

	if len(params) != 3 {
		log.Fatalf("invalid issue title: %s", issue.GetTitle())
	}

	labelNames := make([]string, 0, len(labels))
	for i := 0; i < len(labels); i++ {
		if labels[i].Name == "post" || labels[i].Name == params[1] {
			continue
		}
		labelNames = append(labelNames, labels[i].Name)
	}

	post := Post{
		ID:         fmt.Sprintf("%d", issue.GetID()),
		Title:      params[2],
		Content:    issue.GetBody(),
		Tags:       strings.Join(labelNames, "/"),
		CreateTime: issue.GetCreatedAt().Add(8 * time.Hour).Unix(),
		UpdateTime: issue.GetUpdatedAt().Add(8 * time.Hour).Unix(),
		DisplayId:  params[1],
	}

	_, err = req.
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+token).
		SetBody(post).
		Post("https://blogapi.qianx1.xyz/blog/post")

	if err != nil {
		log.Fatalf("create post failed: %v", err)
	}

	log.Println("create post success")
}
