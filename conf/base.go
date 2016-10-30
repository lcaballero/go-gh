package conf

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// NewClient creates and returns an oauth2 github client based on the
// values provided.
func NewClient(api ApiValues) *github.Client {
	//TODO: check that api.Token is non-nil
	store := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: api.Token})
	authClient := oauth2.NewClient(oauth2.NoContext, store)
	client := github.NewClient(authClient)
	//client.BaseURL = EnterpriseUrl(api.BaseUrl)
	return client
}
