package main

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("MY_GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	// repos, _, err := client.Repositories.List(ctx, "", nil)
	// fmt.Println(repos, err)

	teams, _, err := client.Organizations.ListTeams(ctx, "appscode", nil)
	if err != nil {
		os.Exit(1)
	}
	for _, t := range teams {
		fmt.Println(*t.Name, *t.MembersURL)

		//client.Organizations.list

		mems, _, _ := client.Organizations.ListTeamMembers(ctx, *t.ID, nil)
		for _, m := range mems {
			fmt.Println(*m.ID, *m.Login, m.GetType(), m.GetSiteAdmin())
		}
	}


	//// list all organizations for user "willnorris"
	//orgs, _, err := client.Organizations.List(ctx, "willnorris", nil)
}
