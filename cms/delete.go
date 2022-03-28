package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-github/v43/github"
)

func deletePost() {

	token, err := login()

	if err != nil {
		log.Fatalf("login failed: %v", err)
	}

	cli := github.NewClient(nil)

	issues, _, err := cli.Issues.
		ListByRepo(context.Background(), owner, repo, &github.IssueListByRepoOptions{
			State:  "open",
			Labels: []string{label},
			Sort:   "closed",
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

	_, err = req.
		R().
		SetHeader("Authorization", "Bearer "+token).
		Delete(fmt.Sprintf("https://blogapi.qianx1.xyz/blog/post/%d", issue.GetID()))

	if err != nil {
		log.Fatalf("create post failed: %v", err)
	}

	log.Println("create post success")

}
