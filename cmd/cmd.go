package cmd

import (
	"io"

	"github.com/tanan/wg-in-handy/operator"
	"github.com/urfave/cli/v2"
)

type Command struct {
	Stdout   io.Writer
	Stderr   io.Writer
	Stdin    io.Reader
	Operator *operator.Operator
}

func NewCommand(stdout io.Writer, stderr io.Writer, stdin io.Reader, operator *operator.Operator) *Command {
	return &Command{
		Stdout:   stdout,
		Stderr:   stderr,
		Stdin:    stdin,
		Operator: operator,
	}
}

func (cmd Command) Run(args []string) error {
	var configPath string
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "run as server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "config",
						Usage:       "set config file path",
						Aliases:     []string{"c"},
						Destination: &configPath,
					},
				},
				Action: cmd.runAsAPI,
			},
			{
				Name:  "interface",
				Usage: "manipulate interface",
				Subcommands: []*cli.Command{
					{
						Name:   "show",
						Usage:  "show interface setting",
						Action: cmd.showInterface,
					},
					{
						Name:  "setconf",
						Usage: "add setting via config file",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "config",
								Usage:       "set config file path",
								Aliases:     []string{"c"},
								Destination: &configPath,
							},
						},
						Action: cmd.setConf,
					},
				},
			},
		},
	}
	return app.Run(args)
}

// TODO: implement
func (cmd *Command) runAsAPI(cCtx *cli.Context) error {
	return nil
}

// TODO: implement
func (cmd *Command) showInterface(cCtx *cli.Context) error {
	return nil
}

// TODO: implement
func (cmd *Command) setConf(cCtx *cli.Context) error {
	return nil
}
