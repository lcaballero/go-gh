package conf

// Config represents the structure parsed from the command line
// dictating the required structure/flags.
type Config struct {
	ShowValues    bool   `long:"show-values" description:"Show all values as parsed from command lines and conf files."`
	TokenFile     string `long:"token-file" description:"Name of the file containing the token." default:"~/.go-gh-token"`
	CreateConf    bool   `long:"create-conf" description:"Create bare-bones ~/.go-gh file with guesses for some values."`
	Api           Api    `hidden:"true"`
	ConfFile      string `long:"conf-file" description:"Name of the file where default configuration is stored." default:"~/.go-gh"`
	BaseUrl       string `long:"base-url" description:"Base url to use for rest requests." default:"https://api.github.com/"`
	PR            PR     `command:"pr"`
	Fork          Fork   `command:"fork"`
	Organizations Orgs   `command:"orgs"`
}

// Api holds the configured values for reaching the github api.
type Api struct {
	Active  string
	Current ApiValues
}

// ApiValues hold configured values required for connecting to
// the github api.
type ApiValues struct {
	Token    string
	BaseUrl  string
	Username string
}

// ValidCommand defines how each command should indicate if it is
// truly valid after having been parsed from the command line.
type ValidCommand interface {
	IsValid() bool
}

// NamedCommand provides the name of the command useful for logging
// and debugging.
type NamedCommand interface {
	CmdName() string
}

// A ParsedCommand is both a ValidCommand and NamedCommand as
// defined by those interfaces.
type ParsedCommand interface {
	ValidCommand
	NamedCommand
}
