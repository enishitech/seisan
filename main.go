package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"

	"github.com/enishitech/seisan/config"
	"github.com/enishitech/seisan/expense"
	"github.com/enishitech/seisan/reporter"
	"github.com/enishitech/seisan/request"
)

func main() {
	app := cli.NewApp()
	app.Name = "seisan"
	app.Usage = "Generate seisan report"

	sr := reporter.New(*expense.NewReporter())

	app.Action = func(c *cli.Context) {
		if args := c.Args(); args.Present() {
			conf, err := config.Load("config.yaml")
			if err != nil {
				log.Fatal(err)
			}
			conf.SetTarget(args.First())

			fmt.Printf("Processing %s ...\n", conf.Target)
			reqs, err := request.LoadDir(filepath.Join("data", conf.Target))
			if err != nil {
				log.Fatal(err)
			}
			if err := sr.Report(conf, reqs); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("You must specify the 'TARGET'.\nExample:\n  % seisan 2015/10")
		}
	}
	app.Run(os.Args)
}
