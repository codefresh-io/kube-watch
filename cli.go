package main

import (
	"fmt"
	"os"
	"path/filepath"

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
			Action:      watch,
			Before: func(c *cli.Context) error {
				path := c.String("kube-config")
				dir := os.Getenv("HOME")
				if path[:2] == "~/" {
					path = filepath.Join(dir, path[2:])
					c.Set("kube-config", path)
				}
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "kube-config",
					Value: fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")),
				},
				cli.StringFlag{
					Name:  "url",
					Usage: "Url where to sent the hook",
				},
				cli.StringFlag{
					Name:  "slack-channel-url",
					Usage: "Sent event to slack channel url",
				},
				cli.StringFlag{
					Name:  "watch-type",
					Usage: "Type of event to watch on (Warning, Normal)",
					Value: "ALL",
				},
				cli.StringFlag{
					Name:  "watch-kind",
					Usage: "Kind of K8S resource to watch on (Pod, Service). Alias not supported",
					Value: "ALL",
				},
			},
		},
		{
			Name:        "install",
			Description: "Instal kube-watch on you cluster",
			Action:      installInCluster,
		},
	}
}
