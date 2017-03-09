package client

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/url"
)

func enterpriseURL(baseURL string) *url.URL {
	url, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	return url
}

// NewClient creates and returns an oauth2 github client based on the
// values provided.
func NewClient(api Values) *github.Client {
	//TODO: check that api.Token is non-nil
	store := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: api.Token})
	authClient := oauth2.NewClient(oauth2.NoContext, store)
	client := github.NewClient(authClient)
	client.BaseURL = enterpriseURL(api.BaseUrl)
	return client
}
