package main

import (
	"context"
	"os"

	"golang.org/x/xerrors"

	"github.com/khulnasoft-lab/vul/pkg/commands"
	"github.com/khulnasoft-lab/vul/pkg/log"
	"github.com/khulnasoft-lab/vul/pkg/plugin"

	_ "modernc.org/sqlite" // sqlite driver for RPM DB and Java DB
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Vul behaves as the specified plugin.
	if runAsPlugin := os.Getenv("VUL_RUN_AS_PLUGIN"); runAsPlugin != "" {
		if !plugin.IsPredefined(runAsPlugin) {
			return xerrors.Errorf("unknown plugin: %s", runAsPlugin)
		}
		if err := plugin.RunWithArgs(context.Background(), runAsPlugin, os.Args[1:]); err != nil {
			return xerrors.Errorf("plugin error: %w", err)
		}
		return nil
	}

	app := commands.NewApp()
	if err := app.Execute(); err != nil {
		return err
	}
	return nil
}
