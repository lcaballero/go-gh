package conf

import (
	"errors"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"os"
)

// PR contains the parameters required for a PR POST.
//
// link: https://developer.github.com/v3/pulls/#create-a-pull-request
type PR struct {
	Owner         string
	Repo          string
	Title         string
	Head          string
	Base          string
	Body          string
	Ticket        string
	CurrentBranch string
	ShowHint      bool
	ShowJson      bool
	ShowSummary   bool
	Verbose       bool
	Interactive   bool
}

// IsValid makes sure the correct parameters are non-empty.
func (p PR) IsValid() bool {
	return p.Owner != "" && p.Repo != "" && p.Title != "" && p.Head != "" && p.Base != ""
}

// ApplyConventions
func (p PR) ApplyConventions() (PR, string, error) {
	org, repo, err := CwdConventions(os.Getwd)
	if err != nil {
		return p, "", err
	}

	prefix := p.Prefix()
	p.Owner = org
	p.Repo = repo
	p.Head = fmt.Sprintf("%s:%s", org, p.CurrentBranch)
	p.Body = fmt.Sprintf("%s %s", prefix, p.Title)
	p.Title = fmt.Sprintf("%s %s", prefix, p.Title)

	return p, p.Gist(), nil
}

func (p PR) Prefix() string {
	prefix := ""
	if p.Ticket != "" {
		prefix = fmt.Sprintf("[%s]", p.Ticket)
	}
	if p.CurrentBranch != "" {
		prefix = prefix + fmt.Sprintf("[%s]", p.Base)
	}
	return prefix
}

func (p PR) Gist() string {
	return fmt.Sprintf("PR to merge %s into %s:%s", p.Head, p.Owner, p.Base)
}

// CreatePR issues the post for the command and parameters.
func (p PR) CreatePR(cf Locals) (interface{}, error) {

	client := NewClient(cf.Current)

	pr := &github.NewPullRequest{
		Title: &p.Title,
		Head:  &p.Head,
		Base:  &p.Base,
		Body:  &p.Body,
	}

	ctx := context.Background()
	req, _, err := client.PullRequests.Create(ctx, p.Owner, p.Repo, pr)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (p PR) RunInteractive(cf *Base) (interface{}, error) {
	pr, _, err := p.ApplyConventions()
	if err != nil {
		return nil, err
	}

	fmt.Print("To what branch would you like to have your branch merged into: [%s]", pr.Head)

	return nil, errors.New("not fully implemented yet.")
}
