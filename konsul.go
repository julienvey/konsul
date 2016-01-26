package main

import (
	"os"

	"github.com/jawher/mow.cli"
)

var (
	app       = cli.App("konsul", "CLI to manipulate Konsul")
	consulURL = app.String(cli.StringOpt{
		Name:   "u consul-url",
		Desc:   "URL for Consul Server",
		EnvVar: "CONSUL_URL",
	})
)

func main() {

	app.Command("service", "Manage Consul Services", func(cmd *cli.Cmd) {
		cmd.Command("list", "List Consul Services", listServices)
	})

	app.Run(os.Args)
}
