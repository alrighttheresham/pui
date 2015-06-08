package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pui"
	app.Version = Version
	app.Usage = "This command line application can be used to interact with a PSM server instance"
	app.Author = "Damian ONeill"
	app.Email = "doneill@btisystems.com"
	app.Commands = Commands
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "server, s",
			Value:  "localhost",
			Usage:  "IP Address of the PSM Server",
			EnvVar: "PSM_SERVER",
		}, cli.StringFlag{
			Name:   "username, u",
			Value:  "admin",
			Usage:  "Username for the PSM Server",
			EnvVar: "PSM_USERNAME",
		}, cli.StringFlag{
			Name:   "password, p",
			Value:  "admin",
			Usage:  "Password for the PSM Server",
			EnvVar: "PSM_PASSWORD",
		},
	}
	app.Run(os.Args)
}
