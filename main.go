package main

import (
	"fmt"
	"log"
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
			config, err := loadConfig("config.yaml")
			if err != nil {
				log.Fatal(err)
			}
			config.mergeCliArgs(args)

			fmt.Printf("Processing %s ...\n", config.Target)
			seisanRequests, err := loadSeisanRequests(filepath.Join("data", config.Target))
			if err != nil {
				log.Fatal(err)
			}
			seisanReport := newSeisanReport(seisanRequests, *config)
			if err := seisanReport.export(); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("You must specify the 'TARGET'.\nExample:\n  % seisan 2015/10")
		}
	}
	app.Run(os.Args)
}
