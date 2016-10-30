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
)

// GithubRunnable is a function that will issue github request and
// will return the results or any errors that occurred.
type GithubRunnable func(*conf.Config) (interface{}, error)

// https://github.com/blog/1509-personal-api-tokens
func main() {
	conf := cli.ParseArgs(os.Args[1:]...)

	res, err := run(conf)
	if err != nil {
		panic(err)
	}
	MustShowJSON(res)
}

func run(c *conf.Config) (interface{}, error) {
	var runner GithubRunnable
	name := ""

	if c.Fork.IsValid() {
		name = c.Fork.CmdName()
		runner = c.Fork.CreateFork
	}
	if c.PR.IsValid() {
		name = c.PR.CmdName()
		runner = c.PR.CreatePR
	}

	if c.Organizations.IsValid() {
		name = c.Organizations.CmdName()
		runner = listOrgs
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

func enterpriseURL(baseURL string) *url.URL {
	url, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	return url
}

func showPublicRepos(api conf.ApiValues) {
	client := conf.NewClient(api)
	repos, _, err := client.Repositories.List(api.Username, nil)
	if err != nil {
		panic(err)
	}

	for _, r := range repos {
		fmt.Println(*r.Name)
	}
}

func listOrgs(c *conf.Config) (interface{}, error) {
	client := conf.NewClient(c.Api.Current)

	username := ""
	orgs, _, err := client.Organizations.List(username, nil)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

// MustShowJSON attempts to marshal the given value and panics if an error
// occurs.
func MustShowJSON(e interface{}) {
	bin, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bin))
}
