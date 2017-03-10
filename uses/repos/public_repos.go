package repos

import (
	"fmt"
	"github.com/lcaballero/go-gh/conf"
	gh "github.com/lcaballero/go-gh/uses/client"
	"golang.org/x/net/context"
)

func ShowPublicRepos(locals conf.Locals) {
	rc := locals.Current
	client := gh.NewClient(rc)
	ctx := context.Background()

	repos, _, err := client.Repositories.List(ctx, rc.Username, nil)
	if err != nil {
		panic(err)
	}

	for _, r := range repos {
		fmt.Println(*r.Name)
	}
}
