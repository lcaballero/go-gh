package conf

type Orgs struct {
	List bool `long:"list" description:"Provide a list of user's organizations" required:"true"`
}

func (g Orgs) IsValid() bool {
	return g.List
}
func (g Orgs) CmdName() string {
	return "orgs"
}
