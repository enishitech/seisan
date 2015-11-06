package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "seisan"
	app.Usage = "Generate seisan report"
	app.Action = func(c *cli.Context) {
		if args := c.Args(); args.Present() {
			config := loadConfig("config.yaml")
			config.mergeCliArgs(args)

			fmt.Printf("Processing %s ...\n", config.Target)
			seisanRequests := loadSeisanRequests(filepath.Join("data", config.Target))
			seisanReport := newSeisanReport(seisanRequests, config)
			seisanReport.export()
		} else {
			fmt.Println("You must specify the 'TARGET'.\nExample:\n  % seisan 2015/10")
		}
	}
	app.Run(os.Args)
}
