package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"

	"github.com/enishitech/seisan/config"
)

type Reporter interface {
	Report(*config.Config, []SeisanRequest) error
}

type ExpenseReporter struct{}

func NewExpenseReporter() *ExpenseReporter {
	return &ExpenseReporter{}
}

func (reporter ExpenseReporter) Report(conf *config.Config, requests []SeisanRequest) error {
	seisanReport := newSeisanReport(requests, *conf)
	return seisanReport.export()
}

type SeisanReporter struct {
	reporters []Reporter
}

func NewSeisanReporter(reporters ...Reporter) *SeisanReporter {
	sr := &SeisanReporter{}
	sr.reporters = reporters
	return sr
}

func (sr SeisanReporter) Report(conf *config.Config, requests []SeisanRequest) error {
	for _, r := range sr.reporters {
		err := r.Report(conf, requests)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "seisan"
	app.Usage = "Generate seisan report"

	sr := NewSeisanReporter(*NewExpenseReporter())

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
			if err := sr.Report(conf, seisanRequests); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("You must specify the 'TARGET'.\nExample:\n  % seisan 2015/10")
		}
	}
	app.Run(os.Args)
}
