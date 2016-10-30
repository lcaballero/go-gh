package conf

// Orgs is a command for getting the users Organizations.
type Orgs struct {
	List bool `long:"list" description:"Provide a list of user's organizations" required:"true"`
}

// IsValid returns true if this command is for a list of orgs.
func (g Orgs) IsValid() bool {
	return g.List
}

// CmdName simply returns 'orgs'
func (g Orgs) CmdName() string {
	return "orgs"
}
