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
			Name:   "run",
			Action: onStart, // todo : rename to something like InstallKubeWatchInCluster
		},
		{
			Name:   "start",
			Action: onStart,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "dry-run",
				},
				cli.StringFlag{
					Name:  "kube-config",
					Value: "~/.kube/config",
				},
				cli.StringSliceFlag{
					Name: "watch-on",
				},
			},
		},
	}
}
