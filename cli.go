package main

import (
	"github.com/urfave/cli"
)

func setupCli() *cli.App {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	setupFlags(app)
	return app
}

func setupFlags(app *cli.App) {
	app.Commands = setupCommands()
}

func setupCommands() []cli.Command {
	return []cli.Command{
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
