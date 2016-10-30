package conf

import (
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
)

func NewClient(api ApiValues) *github.Client {
	store := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: api.Token})
	authClient := oauth2.NewClient(oauth2.NoContext, store)
	client := github.NewClient(authClient)
	//client.BaseURL = EnterpriseUrl(api.BaseUrl)

	return client
}

