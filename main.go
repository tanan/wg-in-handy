package main

import (
	"log/slog"
	"os"

	"github.com/tanan/wg-in-handy/cmd"
)

func main() {
	cmd := &cmd.Command{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(cmd.Stdout, nil)))

	if err := cmd.Run(os.Args); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
