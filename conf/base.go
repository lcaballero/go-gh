package conf

// Base represents the structure parsed from the command line
// dictating the required structure/flags.
type Base struct {
	TokenFile          string
	ConfFile           string
	BaseUrl            string
	CreateConf         bool
	ShowValues         bool
	IsUsingConventions bool
}

// Conf holds all the
type Conf struct {
	Base   Base
	Locals Locals
	PR     PR
	Orgs   Orgs
	Fork   Fork
}

// Locals holds the configured values for reaching the github api.
type Locals struct {
	Active  string
	Current Values
}

// Values hold configured values required for connecting to the github api.
type Values struct {
	Token    string
	BaseUrl  string
	Username string
}
