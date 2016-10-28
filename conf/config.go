package conf

type Config struct {
	Action     Fork   `positional-args:"action"`
	TokenFile  string `long:"token-file" description:"Name of the file containing the token." default:"~/.go-gh-token"`
	CreateConf bool   `long:"create-conf" description:"Create bare-bones ~/.go-gh file with guesses for some values."`
	Api        Api
	ConfFile   string `long:"conf-file" description:"Name of the file where default configuration is stored." default:"~/.go-gh"`
	BaseUrl    string `long:"base-url" description:"Base url to use for rest requests." default:"https://github.com/api/v3/"`
}

type Fork struct {
	Owner        string `positional-arg-name:"owner" required:"1"`
	Repo         string `positional-arg-name:"repo" required:"2"`
	Organization string `positional-arg-name:"organization"`
}

type Api struct {
	Active  string
	Current ApiValues
}

type ApiValues struct {
	Token   string
	BaseUrl string
	Username string
}

