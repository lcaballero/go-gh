package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/lcaballero/go-gh/cli"
	"golang.org/x/oauth2"
	"net/url"
	"os"
	"github.com/lcaballero/go-gh/conf"
)

// https://github.com/blog/1509-personal-api-tokens
func main() {
	conf := cli.ParseArgs(os.Args[1:]...)
	MustShowJson(conf)
}

func showPrivateRepos(api conf.ApiValues) {
	store := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: api.Token})
	authClient := oauth2.NewClient(oauth2.NoContext, store)
	client := github.NewClient(authClient)

	repos, _, err := client.Repositories.List(api.Username, &github.RepositoryListOptions{
		Type: "private",
	})
	if err != nil {
		panic(err)
	}

	for _, r := range repos {
		fmt.Println(*r.Name)
	}
}

func EnterpriseUrl(api conf.ApiValues) *url.URL {
	url, err := url.Parse(api.BaseUrl)
	if err != nil {
		panic(err)
	}
	return url
}

func NewClient(api conf.ApiValues) *github.Client {
	store := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: api.Token})
	authClient := oauth2.NewClient(oauth2.NoContext, store)
	client := github.NewClient(authClient)
	client.BaseURL = EnterpriseUrl(api)

	return client
}

func showPublicRepos(api conf.ApiValues) {
	client := NewClient(api)
	repos, _, err := client.Repositories.List(api.Username, nil)
	if err != nil {
		panic(err)
	}

	for _, r := range repos {
		fmt.Println(*r.Name)
	}
}

func createFork(cf *conf.Config) {
	client := NewClient(cf.Api.Current)
	owner, repo := cf.Action.Owner, cf.Action.Repo
	newRepo, _, err := client.Repositories.CreateFork(owner, repo, nil)
	if err != nil {
		panic(err)
	}

	MustShowJson(newRepo)
}

func MustShowJson(e interface{}) {
	bin, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bin))
}
