package pr

import (
	"errors"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/lcaballero/go-gh/conf"
	gh "github.com/lcaballero/go-gh/uses/client"
	"golang.org/x/net/context"
	"os"
	"github.com/lcaballero/go-gh/uses/shared"
)

type PrRequest struct {
	Pr conf.PR
}

// IsValid makes sure the correct parameters are non-empty.
func (r PrRequest) IsValid() bool {
	return r.Pr.Owner != "" && r.Pr.Repo != "" && r.Pr.Title != "" && r.Pr.Head != "" && r.Pr.Base != ""
}

// ApplyConventions
func (r PrRequest) ApplyConventions() (conf.PR, string, error) {
	org, repo, err := shared.CwdConventions(os.Getwd)
	if err != nil {
		return r.Pr, "", err
	}

	prefix := r.Prefix()
	r.Pr.Owner = org
	r.Pr.Repo = repo
	r.Pr.Head = fmt.Sprintf("%s:%s", org, r.Pr.CurrentBranch)
	r.Pr.Body = fmt.Sprintf("%s %s", prefix, r.Pr.Title)
	r.Pr.Title = fmt.Sprintf("%s %s", prefix, r.Pr.Title)

	return r.Pr, r.Gist(), nil
}

func (r PrRequest) Prefix() string {
	prefix := ""
	if r.Pr.Ticket != "" {
		prefix = fmt.Sprintf("[%s]", r.Pr.Ticket)
	}
	if r.Pr.CurrentBranch != "" {
		prefix = prefix + fmt.Sprintf("[%s]", r.Pr.Base)
	}
	return prefix
}

func (r PrRequest) Gist() string {
	p := r.Pr
	return fmt.Sprintf("PR to merge %s into %s:%s", p.Head, p.Owner, p.Base)
}

// CreatePR issues the post for the command and parameters.
func (r PrRequest) CreatePR(cf conf.Locals) (interface{}, error) {

	client := gh.NewClient(cf.Current)

	pr := &github.NewPullRequest{
		Title: &r.Pr.Title,
		Head:  &r.Pr.Head,
		Base:  &r.Pr.Base,
		Body:  &r.Pr.Body,
	}

	ctx := context.Background()
	req, _, err := client.PullRequests.Create(ctx, r.Pr.Owner, r.Pr.Repo, pr)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (p PrRequest) RunInteractive(cf *conf.Base) (interface{}, error) {
	pr, _, err := p.ApplyConventions()
	if err != nil {
		return nil, err
	}

	fmt.Print("To what branch would you like to have your branch merged into: [%s]", pr.Head)

	return nil, errors.New("not fully implemented yet.")
}
