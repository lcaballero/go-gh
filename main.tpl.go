package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"github.com/lcaballero/go-gh/cli"
	"github.com/lcaballero/go-gh/conf"
	. "github.com/lcaballero/go-gh/shared"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"os"
)

// GithubRunnable is a function that will issue github request and
// will return the results or any errors that occurred.
type GithubRunnable func(*conf.Config) (interface{}, error)

// https://github.com/blog/1509-personal-api-tokens
func main() {
	err := cli.ToApp().Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func other() {
	conf := cli.ParseArgs(os.Args[1:]...)

	if conf.ShowValues {
		MustShowJSON(conf)
		return
	}

	res, err := run(conf)
	if err != nil {
		panic(err)
	}

	if res != nil {
		MustShowJSON(res)
	}
}

func run(c *conf.Config) (interface{}, error) {
	var runner GithubRunnable
	var interactive GithubRunnable
	name := ""

	if c.IsUsingConventions {
		pr, gist, err := c.PR.ApplyConventions()
		if err != nil {
			panic(err)
		}
		c.PR = pr

		if !c.PR.ShowJson {
			MustShowJSON(c.PR)
		}

		if !c.PR.ShowHint {
			fmt.Println(gist)
		}

		fmt.Print("Continue? [y/n]: ")
		if !askForConfirmation() {
			return nil, nil
		}
	}

	if c.Fork.IsValid() {
		name = c.Fork.CmdName()
		runner = c.Fork.CreateFork
	}
	if c.PR.IsValid() {
		name = c.PR.CmdName()
		runner = c.PR.CreatePR
		interactive = c.PR.RunInteractive
	}

	if c.Organizations.IsValid() {
		name = c.Organizations.CmdName()
		runner = listOrgs
	}

	if interactive != nil && name != "" {
		return interactive(c)
	}

	if runner != nil && name != "" {
		fmt.Printf("running: %s\n", name)
		return runner(c)
	}

	return nil, fmt.Errorf("couldn't find proper command to run")
}

func showPrivateRepos(api conf.ApiValues) {
	store := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: api.Token})
	authClient := oauth2.NewClient(oauth2.NoContext, store)
	client := github.NewClient(authClient)
	ctx := context.Background()

	repos, _, err := client.Repositories.List(ctx, api.Username, &github.RepositoryListOptions{
		Type: "private",
	})
	if err != nil {
		panic(err)
	}

	for _, r := range repos {
		fmt.Println(*r.Name)
	}
}

func showPublicRepos(api conf.ApiValues) {
	client := conf.NewClient(api)
	ctx := context.Background()

	repos, _, err := client.Repositories.List(ctx, api.Username, nil)
	if err != nil {
		panic(err)
	}

	for _, r := range repos {
		fmt.Println(*r.Name)
	}
}

func listOrgs(c *conf.Config) (interface{}, error) {
	client := conf.NewClient(c.Api.Current)
	ctx := context.Background()

	username := ""
	orgs, _, err := client.Organizations.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}
