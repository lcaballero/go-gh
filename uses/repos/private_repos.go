package repos

import (
	"fmt"
	"github.com/google/go-github/github"
	"github.com/lcaballero/go-gh/conf"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)


func ShowPrivateRepos(locals conf.Locals) {
	api := locals.Current
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

