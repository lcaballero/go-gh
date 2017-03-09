package cli

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/lcaballero/go-gh/conf"
	"io/ioutil"
	"os"
	"strings"
)

const GO_GH = ".go-gh"
const GO_GH_TOKEN = ".go-gh-token"

// Loader loads either the Token file (~/.go-gh-token) or the resources
// file (~/.go-gh).
type Loader struct {
	FileProvider func(string) string
}

// NewLoader creates a Loader with a default FileProvider.
func NewLoader() Loader {
	return Loader{
		FileProvider: func(name string) string {
			home := os.Getenv("HOME")
			file := strings.Replace(name, "~", home, 1)
			return file
		},
	}
}

// LoadToken attempts to find the .go-gh-token file (or it's equivalent)
// and extract the token from within.
func (d Loader) LoadToken(file string) (string, error) {
	if file == "" {
		file = d.FileProvider(GO_GH)
	}

	info, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) {
		return "", nil
	}

	if info.IsDir() {
		return "", fmt.Errorf("token file must be an file not a directory")
	}

	bin, err := ioutil.ReadFile(file)
	if err != nil {
		return "", nil
	}

	s := strings.TrimSpace(string(bin))
	return s, nil
}

// UpdateFromIni expands portions of the config by loading additional
// configuration.
func (d Loader) UpdateFromIni(file string) (conf.Locals, error) {
	locals := conf.Locals{}

	if file == "" {
		file = d.FileProvider(GO_GH_TOKEN)
	}

	info, err := os.Stat(file)
	if err != nil {
		return locals, err
	}

	if info.IsDir() {
		return locals, fmt.Errorf("conf file must be an ini-file not a directory")
	}

	cfg, err := ini.Load(file)
	if err != nil {
		return locals, err
	}

	locals.Active = cfg.Section("").Key("Active").Value()

	s, err := cfg.GetSection(locals.Active)
	if s == nil || err != nil {
		return locals, fmt.Errorf("Could not find section: %s in %s", locals.Active, file)
	}

	err = cfg.Section(locals.Active).MapTo(&locals.Current)
	if err != nil {
		return locals, err
	}

	return locals, nil
}
