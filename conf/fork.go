package conf

// Fork holds the required and optional parameters for issuing a Fork
// request to the github api.
type Fork struct {
	Owner        string
	Repo         string
	Organization string
}

