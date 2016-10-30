package conf

type Config struct {
	TokenFile     string `long:"token-file" description:"Name of the file containing the token." default:"~/.go-gh-token"`
	CreateConf    bool   `long:"create-conf" description:"Create bare-bones ~/.go-gh file with guesses for some values."`
	Api           Api    `hidden:"true"`
	ConfFile      string `long:"conf-file" description:"Name of the file where default configuration is stored." default:"~/.go-gh"`
	BaseUrl       string `long:"base-url" description:"Base url to use for rest requests." default:"https://github.com/api/v3/"`
	PR            PR     `command:"pr"`
	Fork          Fork   `command:"fork"`
	Organizations Orgs   `command:"orgs"`
}

type Api struct {
	Active  string
	Current ApiValues
}

type ApiValues struct {
	Token    string
	BaseUrl  string
	Username string
}

type ValidCommand interface {
	IsValid() bool
}

type NamedCommand interface {
	CmdName() string
}

type ParsedCommand interface {
	ValidCommand
	NamedCommand
}
