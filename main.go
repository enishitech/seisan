package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"

	"github.com/enishitech/seisan/config"
)

func main() {
	app := cli.NewApp()
	app.Name = "seisan"
	app.Usage = "Generate seisan report"
	app.Action = func(c *cli.Context) {
		if args := c.Args(); args.Present() {
			conf, err := config.Load("config.yaml")
			if err != nil {
				log.Fatal(err)
			}
			conf.SetTarget(args.First())

			fmt.Printf("Processing %s ...\n", conf.Target)
			seisanRequests, err := loadSeisanRequests(filepath.Join("data", conf.Target))
			if err != nil {
				log.Fatal(err)
			}
			seisanReport := newSeisanReport(seisanRequests, *conf)
			if err := seisanReport.export(); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("You must specify the 'TARGET'.\nExample:\n  % seisan 2015/10")
		}
	}
	app.Run(os.Args)
}
