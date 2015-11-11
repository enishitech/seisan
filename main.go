package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/tealeg/xlsx"

	"github.com/enishitech/seisan/config"
	"github.com/enishitech/seisan/expense"
	"github.com/enishitech/seisan/request"
)

type Reporter interface {
	Report(*xlsx.Sheet, *config.Config, []request.Request) error
}

type SeisanReporter struct {
	reporters []Reporter
}

func NewSeisanReporter(reporters ...Reporter) *SeisanReporter {
	sr := &SeisanReporter{}
	sr.reporters = reporters
	return sr
}

func renderReportHeader(sheet *xlsx.Sheet, orgName, name string) {
	var row *xlsx.Row
	var cell *xlsx.Cell

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.SetValue(orgName + " 精算シート " + name)
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.SetValue("作成時刻")
	cell = row.AddCell()
	cell.SetValue(time.Now())
	row = sheet.AddRow()
	cell = row.AddCell()
}

func (sr SeisanReporter) Report(conf *config.Config, requests []request.Request) error {
	targetName := strings.Replace(conf.Target, "/", "-", -1)
	xlsx.SetDefaultFont(11, "ＭＳ Ｐゴシック")

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("精算シート")
	if err != nil {
		return err
	}

	renderReportHeader(sheet, targetName, conf.Organization["name"])

	for _, r := range sr.reporters {
		err := r.Report(sheet, conf, requests)
		if err != nil {
			return err
		}
	}

	destPath := filepath.Join("output", targetName+".xlsx")
	if _, err := os.Stat("output"); os.IsNotExist(err) {
		if err := os.Mkdir("output", 0777); err != nil {
			return err
		}
	}
	err = file.Save(destPath)
	if err != nil {
		return err
	}
	fmt.Printf("Wrote to %s\n", destPath)

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "seisan"
	app.Usage = "Generate seisan report"

	sr := NewSeisanReporter(*expense.NewReporter())

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
