package main

import (
	"fmt"
	"github.com/lcaballero/go-gh/cli"
	"github.com/lcaballero/go-gh/conf"
	cmd "gopkg.in/urfave/cli.v2"
	"os"
)

type Runnables struct {
	RunBase func(conf.Base, conf.Locals) error
	RunOrg  func(conf.Orgs, conf.Locals) error
	RunFork func(conf.Fork, conf.Locals) error
	RunPr   func(conf.PR, conf.Locals) error
}

// https://github.com/blog/1509-personal-api-tokens
func main() {
	runners := Runnables{
		RunBase: func(base conf.Base, api conf.Locals) error {
			fmt.Println("run base")
			return nil
		},
		RunOrg: func(orgs conf.Orgs, api conf.Locals) error {
			fmt.Println("run orgs")
			return nil
		},
		RunFork: func(fork conf.Fork, api conf.Locals) error {
			fmt.Println("run fork")
			return nil
		},
		RunPr: func(pr conf.PR, api conf.Locals) error {
			fmt.Println("run pr")
			return nil
		},
	}

	proc := cli.Processing{
		BeforeAction: func(ctx *cmd.Context) error {
			fmt.Println("before action")
			return nil
		},
		BaseAction: func(ctx *cmd.Context) error {
			fmt.Println("base action")
			runners.RunBase(conf.Base{}, conf.Locals{})
			return nil
		},
		OrgAction: func(ctx *cmd.Context) error {
			fmt.Println("org action")
			runners.RunOrg(conf.Orgs{}, conf.Locals{})
			return nil
		},
		ForkAction: func(ctx *cmd.Context) error {
			fmt.Println("fork action")
			runners.RunFork(conf.Fork{}, conf.Locals{})
			return nil
		},
		PrAction: func(ctx *cmd.Context) error {
			fmt.Println("pr action")
			runners.RunPr(conf.PR{}, conf.Locals{})
			return nil
		},
	}

	err := cli.New(proc).Run(os.Args)
	if err != nil {
		panic(err)
	}
}
