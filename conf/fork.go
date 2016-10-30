package conf

// Fork holds the required and optional parameters for issuing a Fork
// request to the github api.
type Fork struct {
	Owner        string `long:"owner" short:"o" description:"Owner of the branch to fork" required:"1"`
	Repo         string `long:"repo" short:"r" description:"Name of the repo to fork" required:"2"`
	Organization string `long:"org" description:"Name of the organization to fork into"`
}

// CmdName simply returns "fork".
func (f Fork) CmdName() string {
	return "fork"
}

// IsValid checks that the required parameters are non-empty strings.
func (f Fork) IsValid() bool {
	return f.Owner != "" && f.Repo != ""
}

// CreateFork issues the create fork request to github api.
func (f Fork) CreateFork(cf *Config) (interface{}, error) {

	//TODO: check again for command validity.

	client := NewClient(cf.Api.Current)
	owner, repo := f.Owner, f.Repo

	newRepo, _, err := client.Repositories.CreateFork(owner, repo, nil)
	if err != nil {
		return nil, err
	}
	return newRepo, nil
}
