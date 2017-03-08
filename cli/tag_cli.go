package cli

import (
	"fmt"
	"github.com/lcaballero/go-gh/conf"
	"gopkg.in/urfave/cli.v2"
	"reflect"
	"strconv"
)

type FlagStruct struct {
	Name   string
	Usage  string
	Action cli.ActionFunc
	Struct interface{}
}

var commands = []FlagStruct{
	FlagStruct{
		Name:  "pr",
		Usage: "Create a PR",
		Action: func(ctx *cli.Context) error {
			fmt.Println("Creating a PR")
			return nil
		},
		Struct: conf.PR{},
	},
	FlagStruct{
		Name:  "fork",
		Usage: "Create a fork",
		Action: func(ctx *cli.Context) error {
			fmt.Println("Creating a fork")
			return nil
		},
		Struct: conf.Fork{},
	},
	FlagStruct{
		Name: "org",
		Usage: "Create a Org",
		Action: func(ctx *cli.Context) error {
			fmt.Println("Create a Org")
			return nil
		},
		Struct: conf.Orgs{},
	},
}

func ToApp() *cli.App {
	app := &cli.App{
		Name:    "go-gh",
		Version: "0.0.1",
		Usage:   "A CLI to interface with github from inside a repository directory",
	}

	cmds := make([]*cli.Command, 0)

	for _,cmd := range commands {
		c := ToCommand(cmd)
		cmds = append(cmds, c)
	}

	app.Commands = cmds

	return app
}

func ToCommand(c FlagStruct) *cli.Command {
	val := reflect.ValueOf(c.Struct)
	ty := val.Type()
	fn := ty.NumField()

	flags := make([]cli.Flag, 0)

	for i := 0; i < fn; i++ {
		field := ty.Field(i)
		long := field.Tag.Get("long")
		desc := field.Tag.Get("description")
		def := field.Tag.Get("default")
		kind := field.Type.Kind()

		f := ToFlag(kind, long, desc, def)

		if f != nil {
			flags = append(flags, f)
		}
	}

	cmd := &cli.Command{
		Name:   c.Name,
		Usage:  c.Usage,
		Action: c.Action,
		Flags:  flags,
	}

	return cmd
}

func ToFlag(kind reflect.Kind, long, desc, def string) cli.Flag {
	switch kind {
	case reflect.Bool:
		var sf cli.Flag = &cli.BoolFlag{
			Name:  long,
			Usage: desc,
			Value: ToBool(def),
		}
		return sf
	case reflect.Int:
		var sf cli.Flag = &cli.IntFlag{
			Name:  long,
			Usage: desc,
			Value: ToInt(def),
		}
		return sf
	case reflect.String:
		var sf cli.Flag = &cli.StringFlag{
			Name:  long,
			Usage: desc,
			Value: def,
		}
		return sf
	default:
		return nil
	}
}

func ToBool(s string) bool {
	switch s {
	case "true":
		return true
	default:
		return false
	}
}

func ToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}