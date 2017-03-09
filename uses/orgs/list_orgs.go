package orgs

import (
	"github.com/lcaballero/go-gh/conf"
	gh "github.com/lcaballero/go-gh/uses/client"
	"golang.org/x/net/context"
)

func ListOrgs(c conf.Locals) (interface{}, error) {
	client := gh.NewClient(c.Current)
	ctx := context.Background()

	username := ""
	orgs, _, err := client.Organizations.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}
