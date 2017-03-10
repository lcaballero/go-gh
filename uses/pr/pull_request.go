package pr

import (
	"errors"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/lcaballero/go-gh/conf"
	app "github.com/lcaballero/go-gh/shared"
	gh "github.com/lcaballero/go-gh/uses/client"
	"github.com/lcaballero/go-gh/uses/shared"
	"golang.org/x/net/context"
	"os"
	"strings"
	"github.com/lcaballero/go-gh/uses/resources"
)

const WS = " \n\r\t"

// PrRequest encapsulates the information to make an API request for making
// a pull-request.
type PrRequest struct {
	Pr     *conf.PR
	Base   *conf.Base
	Locals *conf.Locals
	Loader resources.Loader
}

// New loads required values from the context to produce the Request.
func New(vals conf.ValueContext) *PrRequest {
	pr := conf.LoadPr(vals)
	base := conf.LoadBase(vals)
	return &PrRequest{
		Pr:     &pr,
		Base:   &base,
		Loader: resources.New(),
	}
}

// IsValid makes sure the correct parameters are non-empty.
func (r *PrRequest) IsValid() bool {
	return r.Pr.Owner != "" && r.Pr.Repo != "" &&
		r.Pr.Title != "" && r.Pr.Head != "" &&
		r.Pr.Base != "" && r.Pr.CurrentBranch != ""
}

func (r *PrRequest) ApplyBranch() error {
	branch, err := app.CurrentBranch()
	if err != nil {
		return nil
	}
	r.Pr.CurrentBranch = branch
	return nil
}

func (r *PrRequest) ApplyLocals() error {
	locals, err := r.Loader.LoadLocals()
	if err != nil {
		fmt.Println(err)
		return err
	}
	r.Locals = &locals
	return nil
}

// ApplyConventions
func (r *PrRequest) ApplyConventions() error {
	org, repo, err := shared.CwdConventions(os.Getwd)
	if err != nil {
		return err
	}

	prefix := r.Prefix()

	r.Pr.Owner = org
	r.Pr.Repo = repo
	r.Pr.Head = fmt.Sprintf("%s:%s", org, r.Pr.CurrentBranch)
	r.Pr.Body = strings.Trim(fmt.Sprintf("%s %s", prefix, r.Pr.Title), WS)
	r.Pr.Title = strings.Trim(fmt.Sprintf("%s %s", prefix, r.Pr.Title), WS)

	return nil
}

func (r *PrRequest) Prefix() string {
	prefix := ""
	if r.Pr.Ticket != "" {
		prefix = fmt.Sprintf("[%s]", r.Pr.Ticket)
	}
	if r.Pr.CurrentBranch != "" {
		prefix = prefix + fmt.Sprintf("[%s]", r.Pr.Base)
	}
	return prefix
}

func (r *PrRequest) Gist() string {
	return fmt.Sprintf("PR to merge %s into %s:%s",
		r.Pr.Head, r.Pr.Owner, r.Pr.Base)
}

// CreatePR issues the post for the command and parameters.
func (r *PrRequest) CreatePR() (interface{}, error) {
	if !r.IsValid() {
		return nil, errors.New("could not validate the request propertly")
	}

	client := gh.NewClient(r.Locals.Current)

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

func (p *PrRequest) RunInteractive(cf *conf.Base) error {
	err := p.ApplyConventions()
	if err != nil {
		return err
	}

	fmt.Print("To what branch would you like your branch merged into: [%s]",
		p.Pr.Head)

	return errors.New("not fully implemented yet.")
}
