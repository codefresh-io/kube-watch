package main

import (
	"os"
)

func main() {
	app := setupCli()
	// app.Action = runDryRun // just for debug the runDryRun
	app.Run(os.Args)
}
