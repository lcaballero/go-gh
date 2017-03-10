package cli

import (
	"gopkg.in/urfave/cli.v2"
)

type Processing struct {
	BaseAction   cli.ActionFunc
	OrgAction    cli.ActionFunc
	ForkAction   cli.ActionFunc
	PrAction     cli.ActionFunc
}

func New(proc Processing) *cli.App {
	app := &cli.App{
		Name:    "go-gh",
		Version: "0.0.1",
		Usage:   "A CLI to interface with github from inside a repository directory",
		Action:  proc.BaseAction,
		Flags:   BaseFlags(),
		Commands: []*cli.Command{
			PrCommand(proc.PrAction),
			ForkCommand(proc.ForkAction),
			OrgCommand(proc.OrgAction),
		},
	}
	return app
}

func OrgCommand(action cli.ActionFunc) *cli.Command {
	return &cli.Command{
		Name:   "org",
		Usage:  "Manipulate organizations from within a repository",
		Action: action,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "list",
				Usage: "Provide a list of user's organizations",
			},
		},
	}
}

func ForkCommand(action cli.ActionFunc) *cli.Command {
	return &cli.Command{
		Name:   "fork",
		Usage:  "Create fork from within a repository",
		Action: action,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "owner",
				Usage: "Owner of the branch to fork",
			},
			&cli.StringFlag{
				Name:  "repo",
				Usage: "Name of the repo to fork",
			},
			&cli.StringFlag{
				Name:  "org",
				Usage: "Name of the organization to fork into",
			},
		},
	}
}

func PrCommand(action cli.ActionFunc) *cli.Command {
	return &cli.Command{
		Name:   "pr",
		Usage:  "Create pull-request from within a repository",
		Action: action,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "owner",
				Usage: "Owner of the destination repo to pull changes into.",
			},
			&cli.StringFlag{
				Name:  "repo",
				Usage: "Source repository to pull changes from.",
			},
			&cli.StringFlag{
				Name:  "title",
				Usage: "The title of the pull request",
			},
			&cli.StringFlag{
				Name:  "head",
				Usage: "The name of the branch where the changes are implemented. For cross repository pull requests in the same network, namespace head with a user like this: 'username:branch'",
			},
			&cli.StringFlag{
				Name:  "base",
				Usage: "The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository.",
				Value: "master",
			},
			&cli.StringFlag{
				Name:  "body",
				Usage: "The contents of the pull request.",
			},
			&cli.StringFlag{
				Name:  "ticket",
				Usage: "A ticket number associated with the PR.",
			},
			&cli.StringFlag{
				Name:  "current-branch",
				Usage: "Current branch -- until this can be derived via git command line tool.",
			},
			&cli.BoolFlag{
				Name:  "show-hint",
				Usage: "Hide the display of the gist, which shows source-branch into dest-branch text.",
			},
			&cli.BoolFlag{
				Name:  "show-json",
				Usage: "Hides the json post content.",
			},
			&cli.BoolFlag{
				Name:  "show-summary",
				Usage: "Show non-json human readable summary.",
			},
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "Display additional output produced when creating PR parameters to the API.",
			},
			&cli.BoolFlag{
				Name:  "interactive",
				Usage: "Causes the utility to ask question to fill in the parameters rather than using flags.",
			},
		},
	}
}

func BaseFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "show-values",
			Usage: "Show all values as parsed from command lines and conf files then take no action.",
		},
		&cli.StringFlag{
			Name:  "token-file",
			Usage: "Name of the file containing the token.",
			Value: "~/.go-gh-token",
		},
		&cli.StringFlag{
			Name:  "create-conf",
			Usage: "Create bare-bones ~/.go-gh file with guesses for some values.",
		},
		&cli.StringFlag{
			Name:  "conf-file",
			Usage: "Name of the file where default configuration is stored.",
			Value: "~/.go-gh",
		},
		&cli.StringFlag{
			Name:  "base-url",
			Usage: "Base url to use for rest requests.",
			Value: "https://api.github.com/",
		},
		&cli.BoolFlag{
			Name:  "using-convention",
			Usage: "Use PWD conventions from /[git.host]/[organization]/[repo] to populate parameters.",
		},
	}
}
