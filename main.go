package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"

	"github.com/enishitech/seisan/expense"
	"github.com/enishitech/seisan/reporter"
)

func main() {
	app := cli.NewApp()
	app.Name = "seisan"
	app.Usage = "Generate seisan report"

	sr := reporter.New(*expense.NewReporter())

	app.Action = func(c *cli.Context) {
		if args := c.Args(); args.Present() {
			target := args.First()

			fmt.Printf("Processing %s ...\n", target)
			if err := sr.Report(".", target); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("You must specify the 'TARGET'.\nExample:\n  % seisan 2015/10")
		}
	}
	app.Run(os.Args)
}
