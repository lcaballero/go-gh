package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/lcaballero/go-gh/cli"
	"github.com/lcaballero/go-gh/conf"
	"golang.org/x/oauth2"
	"net/url"
	"os"
	"github.com/lcaballero/griller/cmd"
)

type GithubRunnable func(*conf.Config) (interface{}, error)

// https://github.com/blog/1509-personal-api-tokens
func main() {
	conf := cli.ParseArgs(os.Args[1:]...)

	res, err := run(conf)
	if err != nil {
		panic(err)
	}
	MustShowJson(res)
}

func run(c *conf.Config) (interface{}, error) {
	var runner GithubRunnable = nil
	name := ""

	if c.Fork.IsValid() {
		name = c.Fork.CmdName()
		runner = c.Fork.CreateFork
	}
	if c.PR.IsValid() {
		name = c.PR.CmdName()
		runner = CreatePR
	}

	if c.Organizations.IsValid() {
		name = c.Organizations.CmdName()
		runner = ListOrgs
	}

	if runner != nil && name != "" {
		fmt.Printf("running: %s\n", name)
		return runner(c)
	}

	return nil, fmt.Errorf("coldn't find proper command to run")
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

func EnterpriseUrl(baseUrl string) *url.URL {
	url, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}
	return url
}

func NewClient(api conf.ApiValues) *github.Client {
	store := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: api.Token})
	authClient := oauth2.NewClient(oauth2.NoContext, store)
	client := github.NewClient(authClient)
	//client.BaseURL = EnterpriseUrl(api.BaseUrl)

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

func CreatePR(cf *conf.Config) (interface{}, error) {
	return nil, nil
}

func ListOrgs(c *conf.Config) (interface{}, error) {
	client := NewClient(c.Api.Current)

//	username := c.Api.Current.Username
	username := ""
	orgs, _, err := client.Organizations.List(username, nil)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

func MustShowJson(e interface{}) {
	bin, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bin))
}
