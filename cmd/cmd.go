package cmd

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/tanan/wg-in-handy/api"
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
				Name:  "api",
				Usage: "run as api",
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
	router := api.NewRouter()
	router.Run(":8080")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("error when listen", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown Server ...")

	return nil
}

// TODO: implement
func (cmd *Command) showInterface(cCtx *cli.Context) error {
	inf := cmd.Operator.ShowInterface()
	slog.Info("interface", slog.String("Name", inf.Name), slog.String("Address", inf.Address))
	return nil
}

// TODO: implement
func (cmd *Command) setConf(cCtx *cli.Context) error {
	return nil
}
