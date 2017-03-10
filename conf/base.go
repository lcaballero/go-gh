package conf

import (
	"encoding/json"
)

// Base represents the structure parsed from the command line
// dictating the required structure/flags.
type Base struct {
	TokenFile        string
	ConfFile         string
	BaseUrl          string
	CreateConf       bool
	ShowValues       bool
	UsingConventions bool
}

func LoadBase(ctx ValueContext) Base {
	c := ContextLoader{ctx}
	base := Base{}
	c.String("token-file", &base.TokenFile)
	c.String("conf-file", &base.ConfFile)
	c.String("base-url", &base.BaseUrl)
	c.Bool("create-conf", &base.CreateConf)
	c.Bool("show-values", &base.ShowValues)
	c.Bool("using-convention", &base.UsingConventions)
	return base
}

func (pr Base) ToJson() string {
	bin, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(bin)
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

func (pr Locals) ToJson() string {
	bin, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(bin)
}