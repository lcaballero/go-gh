package main

import (
	"fmt"
	"github.com/lcaballero/go-gh/cli"
	pr_req "github.com/lcaballero/go-gh/uses/pr"
	cmd "gopkg.in/urfave/cli.v2"
	"os"
	"encoding/json"
)

// https://github.com/blog/1509-personal-api-tokens
func main() {
	proc := cli.Processing{
		BaseAction: func(ctx *cmd.Context) error {
			fmt.Println("base action")
			return nil
		},
		OrgAction: func(ctx *cmd.Context) error {
			fmt.Println("org action")
			return nil
		},
		ForkAction: func(ctx *cmd.Context) error {
			fmt.Println("fork action")
			return nil
		},
		PrAction: func(ctx *cmd.Context) error {
			req := pr_req.New(ctx)

			err := req.ApplyBranch()
			if err != nil {
				return err
			}

			err = req.ApplyLocals()
			if err != nil {
				return err
			}

			err = req.ApplyConventions()
			if err != nil {
				return err
			}

			fmt.Println("pr:", req.Pr.ToJson())
			fmt.Println("base:", req.Base.ToJson())
			fmt.Println("locals:", req.Locals.ToJson())
			fmt.Println(req.Gist())
			fmt.Println("is valid:", req.IsValid())

			res, err := req.CreatePR()
			if err != nil {
				return err
			}

			bin, err := json.MarshalIndent(res, "", "  ")
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("result:")
			fmt.Println(string(bin))

			return nil
		},
	}

	err := cli.New(proc).Run(os.Args)
	if err != nil {
		panic(err)
	}
}
