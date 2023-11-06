package main

import (
	"os"

	"github.com/aquasecurity/vul/pkg/commands"
	"github.com/aquasecurity/vul/pkg/log"
)

var (
	version = "dev"
)

func main() {
	app := commands.NewApp(version)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
