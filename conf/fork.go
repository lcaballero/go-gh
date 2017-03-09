package conf

import (
	"fmt"
	"golang.org/x/net/context"
)

// Fork holds the required and optional parameters for issuing a Fork
// request to the github api.
type Fork struct {
	Owner        string
	Repo         string
	Organization string
}

// IsValid checks that the required parameters are non-empty strings.
func (f Fork) IsValid() bool {
	return f.Owner != "" && f.Repo != ""
}

// CreateFork issues the create fork request to github api.
func (f Fork) CreateFork(cf Locals) (interface{}, error) {
	if f.IsValid() {
		return nil, fmt.Errorf("fork doesn't have required parameters")
	}

	client := NewClient(cf.Current)
	owner, repo := f.Owner, f.Repo

	ctx := context.Background()
	newRepo, _, err := client.Repositories.CreateFork(ctx, owner, repo, nil)
	if err != nil {
		return nil, err
	}
	return newRepo, nil
}
