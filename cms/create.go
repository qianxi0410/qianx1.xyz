package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-github/v43/github"
)

func createPost() {
	token, err := login()

	if err != nil {
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

	labelNames := make([]string, 0, len(labels)-1)
	for i := 1; i < len(labels); i++ {
		labelNames = append(labelNames, labels[i].Name)
	}

	post := Post{
		ID:      fmt.Sprintf("%d", issue.GetID()),
		Title:   issue.GetTitle(),
		Content: issue.GetBody(),
		Tags:    strings.Join(labelNames, "/"),
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
