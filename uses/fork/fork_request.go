package fork

import (
	"fmt"
	"github.com/lcaballero/go-gh/conf"
	gh "github.com/lcaballero/go-gh/uses/client"
	"golang.org/x/net/context"
)

type ForkRequest struct {
	Fork conf.Fork
}

// IsValid checks that the required parameters are non-empty strings.
func (f ForkRequest) IsValid() bool {
	return f.Fork.Owner != "" && f.Fork.Repo != ""
}

// CreateFork issues the create fork request to github api.
func (f ForkRequest) CreateFork(cf conf.Locals) (interface{}, error) {
	if f.IsValid() {
		return nil, fmt.Errorf("fork doesn't have required parameters")
	}

	client := gh.NewClient(cf.Current)
	owner, repo := f.Fork.Owner, f.Fork.Repo

	ctx := context.Background()
	newRepo, _, err := client.Repositories.CreateFork(ctx, owner, repo, nil)
	if err != nil {
		return nil, err
	}
	return newRepo, nil
}
