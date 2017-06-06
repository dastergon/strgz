package lib

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func ConnectGithub(token string) *github.Client {
	client := github.NewClient(nil)
	if token != "" {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}
	return client
}
