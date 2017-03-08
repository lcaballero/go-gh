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
	Owner         string `long:"owner" description:"Owner of the destination repo to pull changes into."`
	Repo          string `long:"repo" description:"Source repository to pull changes from."`
	Title         string `long:"title" description:"The title of the pull request"`
	Head          string `long:"head" description:"The name of the branch where the changes are implemented. For cross repository pull requests in the same network, namespace head with a user like this: 'username:branch'"`
	Base          string `long:"base" description:"The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository." default:"master"`
	Body          string `long:"body" description:"The contents of the pull request."`
	Ticket        string `long:"ticket" description:"A ticket number associated with the PR."`
	CurrentBranch string `long:"branch" description:"Current branch -- until this can be derived via git command line tool."`
	ShowHint      bool   `long:"show-hint" description:"Hide the display of the gist, which shows source-branch into dest-branch text." default:"true"`
	ShowJson      bool   `long:"show-json" description:"Hides the json post content." default:"true"`
	ShowSummary   bool   `long:"show-json" description:"Show non-json human readable summary." default:"true"`
	Verbose       bool   `long:"verbose" description:"Display additional output produced when creating PR parameters to the API."`
	Interactive   bool   `long:"interactive" description:"Causes the utility to ask question to fill in the parameters rather than using flags."`
}

// IsValid makes sure the correct parameters are non-empty.
func (p PR) IsValid() bool {
	return p.Owner != "" && p.Repo != "" && p.Title != "" && p.Head != "" && p.Base != ""
}

// CmdName returns the name of the command.
func (p PR) CmdName() string {
	return "pr"
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
func (p PR) CreatePR(cf *Config) (interface{}, error) {

	//TODO: check/re-check the command is valid as per the method above.

	client := NewClient(cf.Api.Current)

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

func (p PR) RunInteractive(cf *Config) (interface{}, error) {
	pr, _, err := p.ApplyConventions()
	if err != nil {
		return nil, err
	}

	fmt.Print("To what branch would you like to have your branch merged into: [%s]", pr.Head)

	return nil, errors.New("not fully implemented yet.")
}
