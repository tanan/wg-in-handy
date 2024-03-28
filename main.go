package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/tanan/wg-in-handy/cmd"
	"github.com/tanan/wg-in-handy/config"
)

func main() {
	cmd := &cmd.Command{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(cmd.Stdout, nil)))

	cfg := config.NewConfig()
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	if err := cmd.Run(os.Args); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
