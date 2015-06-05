package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pui"
	app.Version = Version
	app.Usage = ""
	app.Author = "Damian ONeill"
	app.Email = "doneill@btisystems.com"
	app.Commands = Commands

	app.Run(os.Args)
}
