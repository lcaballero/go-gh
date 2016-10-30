package conf

import "github.com/google/go-github/github"

// PR contains the parameters required for a PR POST.
//
// link: https://developer.github.com/v3/pulls/#create-a-pull-request
type PR struct {
	Owner string `long:"owner" short:"o" description:"Target owner of repo to pull changes into." required:"true"`
	Repo  string `long:"repo" short:"r" description:"Target repository to pull changes into." required:"true"`
	Title string `long:"title" short:"t" description:"The title of the pull request" required:"true"`
	Head  string `long:"head" description:"The name of the branch where the changes are implemented. For cross repository pull requests in the same network, namespace head with a user like this: 'username:branch'" required:"true"`
	Base  string `long:"base" short:"b" description:"The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository." required:"true"`
	Body  string `long:"body" description:"The contents of the pull request."`
}

// IsValid makes sure the correct parameters are non-empty.
func (p PR) IsValid() bool {
	return p.Owner != "" && p.Repo != "" && p.Title != "" && p.Head != "" && p.Base != ""
}

// CmdName returns the name of the command.
func (p PR) CmdName() string {
	return "pr"
}

// CreatePR issues the post for the command and parameters.
func (p PR) CreatePR(cf *Config) (interface{}, error) {

	//TODO: check/re-check the command is valid as per the method above.

	client := NewClient(cf.Api.Current)

	owner, repo := p.Owner, p.Repo
	pr := &github.NewPullRequest{
		Title: &p.Title,
		Head:  &p.Head,
		Base:  &p.Base,
		Body:  &p.Body,
	}

	req, _, err := client.PullRequests.Create(owner, repo, pr)
	if err != nil {
		return nil, err
	}

	return req, nil
}
