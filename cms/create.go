package main

import (
	"context"
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

	issue := issues[0]

	reg := regexp.MustCompile("\\[([a-zA-Z0-9-_!]+)\\] ([\u4e00-\u9fa5_a-zA-Z0-9!\\s]+)")
	params := reg.FindStringSubmatch(issue.GetTitle())

	if len(params) != 3 {
		log.Fatalf("invalid issue title: %s", issue.GetTitle())
	}

	labelNames := make([]string, 0, len(issue.Labels))
	flag := false
	for _, label := range issue.Labels {
		if label.GetName() == "post" {
			flag = true
			continue
		}
		labelNames = append(labelNames, label.GetName())
	}

	if !flag {
		log.Println("no post label found, so don't create post")
		return
	}

	post := Post{
		ID:          fmt.Sprintf("%d", issue.GetID()),
		Title:       params[2],
		Content:     issue.GetBody(),
		Tags:        strings.Join(labelNames, "/"),
		CreateTime:  issue.GetCreatedAt().Add(8 * time.Hour).Unix(),
		UpdateTime:  issue.GetUpdatedAt().Add(8 * time.Hour).Unix(),
		DisplayId:   params[1],
		IssueNumber: issue.GetNumber(),
	}

	req := resty.New()
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
