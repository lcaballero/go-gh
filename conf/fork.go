package conf

import "encoding/json"

// Fork holds the required and optional parameters for issuing a Fork
// request to the github api.
type Fork struct {
	Owner        string
	Repo         string
	Organization string
}

func LoadFork(ctx ValueContext) Fork {
	c := ContextLoader{ctx}
	fork := Fork{}
	c.String("owner", &fork.Owner)
	c.String("organization", &fork.Organization)
	c.String("repo", &fork.Repo)
	return fork
}

func (fork Fork) ToJson() string {
	bin, err := json.MarshalIndent(fork, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(bin)
}
