package main

import (
	"github.com/urfave/cli"
)

func setupCli() *cli.App {
	app := cli.NewApp()
	app.Name = "kube-watch"
	setupCommands(app)
	return app
}

func setupCommands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:        "run",
			Description: "Watch from localhost on current-context in ~/.kube/config",
			Action:      dryRun,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "kube-config",
					Value: "~/.kube/config",
				},
			},
		},
	}
}
