package cli

import (
	"fmt"
	"github.com/go-ini/ini"
	cmd "github.com/codegangsta/cli"
	"github.com/lcaballero/go-gh/conf"
	"io/ioutil"
	"os"
	"strings"
)

// ParseArgs parses the command lines args (minus the first that is
// the command go-gh itself).
func ParseArgs(args ...string) *conf.Config {
	cfg := &conf.Config{}
	parser := flags.NewParser(cfg, flags.Default)
	_, err := parser.ParseArgs(args)
	if err != nil {
		os.Exit(1)
	}

	cfg.Api.Current = conf.ApiValues{
		BaseUrl: cfg.BaseUrl,
	}

	err = loadToken(cfg)
	if err != nil {
		parser.WriteHelp(os.Stderr)
		os.Exit(1)
	}

	err = updateFromIni(cfg)
	if os.IsNotExist(err) {
		return cfg
	}

	if err != nil {
		fmt.Printf("Failed to parse %s\n%s\n\n", cfg.ConfFile, err)
		parser.WriteHelp(os.Stderr)
		os.Exit(1)
	}

	return cfg
}

// loadToken attempts to find the .go-gh-token file (or it's equivalent)
// and extract the token from within.
func loadToken(c *conf.Config) error {
	if c.TokenFile == "" {
		return nil
	}
	home := os.Getenv("HOME")
	file := strings.Replace(c.TokenFile, "~", home, 1)

	info, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) {
		return nil
	}
	if info.IsDir() {
		return fmt.Errorf("token file must be an file not a directory")
	}

	bin, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	s := string(bin)
	c.Api.Current.Token = strings.TrimSpace(s)
	return nil
}

// updateFromIni expands portions of the config by loading additional
// configuration.
func updateFromIni(c *conf.Config) error {
	if c.ConfFile == "" {
		return nil
	}

	home := os.Getenv("HOME")
	file := strings.Replace(c.ConfFile, "~", home, 1)

	info, err := os.Stat(file)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("conf file must be an ini-file not a directory")
	}

	cfg, err := ini.Load(file)
	if err != nil {
		return err
	}

	//TODO: make this case-insensitive at some point.
	active := cfg.Section("").Key("Active").Value()

	//TODO: make this case-insensitive at some point.
	s, err := cfg.GetSection(active)
	if s == nil || err != nil {
		return fmt.Errorf("Could not find section: %s in %s", active, c.ConfFile)
	}

	err = cfg.Section(active).MapTo(&c.Api.Current)
	if err != nil {
		fmt.Println(err)
	}

	c.Api.Active = active

	return nil
}
