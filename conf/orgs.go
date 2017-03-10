package conf

// Orgs is a command for getting the users Organizations.
type Orgs struct {
	List bool
}

// IsValid returns true if this command is for a list of orgs.
func (g Orgs) IsValid() bool {
	return g.List
}

func LoadOrgs(ctx ValueContext) Orgs {
	c := ContextLoader{ctx}
	orgs := Orgs{}
	c.Bool("list", &orgs.List)
	return orgs
}
