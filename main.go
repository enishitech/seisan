package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
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

type ExpenseReporter struct{}

func NewExpenseReporter() *ExpenseReporter {
	return &ExpenseReporter{}
}

type ExpenseRequest struct {
	Applicant       string          `yaml:"applicant"`
	ExpeneseEntries []expense.Entry `yaml:"expense"`
}

func (reporter ExpenseReporter) renderSummary(sheet *xlsx.Sheet, sumByApplicant map[string]int) {
	var row *xlsx.Row
	var cell *xlsx.Cell

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.SetValue("立替払サマリー")
	row = sheet.AddRow()
	for _, heading := range []string{"氏名", "金額"} {
		cell = row.AddCell()
		cell.SetValue(heading)
	}
	for key, value := range sumByApplicant {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.SetValue(key)
		cell = row.AddCell()
		cell.SetValue(value)
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.SetValue("")
}

type ByDate []expense.Entry

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date < a[j].Date }

func (reporter ExpenseReporter) renderEntries(sheet *xlsx.Sheet, entries []expense.Entry) {
	var row *xlsx.Row
	var cell *xlsx.Cell

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.SetValue("立替払明細")
	row = sheet.AddRow()
	for _, heading := range []string{"日付", "立替者", "金額", "摘要", "備考"} {
		cell = row.AddCell()
		cell.SetValue(heading)
	}
	for _, detail := range entries {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.SetValue(detail.Date)
		cell = row.AddCell()
		cell.SetValue(detail.Applicant)
		cell = row.AddCell()
		cell.SetValue(detail.Amount)
		cell = row.AddCell()
		cell.SetValue(detail.Remarks)
	}
}

func (reporter ExpenseReporter) Report(sheet *xlsx.Sheet, conf *config.Config, requests []request.Request) error {
	entries := make([]expense.Entry, 0)
	sumByApplicant := make(map[string]int)
	for _, req := range requests {
		var er ExpenseRequest
		if err := req.Unmarshal(&er); err != nil {
			return err
		}
		if _, ok := sumByApplicant[er.Applicant]; !ok {
			sumByApplicant[er.Applicant] = 0
		}
		for _, entry := range er.ExpeneseEntries {
			entry.Applicant = er.Applicant
			sumByApplicant[er.Applicant] += entry.Amount
			entries = append(entries, entry)
		}
	}
	sort.Sort(ByDate(entries))

	reporter.renderSummary(sheet, sumByApplicant)
	reporter.renderEntries(sheet, entries)

	return nil
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

	sr := NewSeisanReporter(*NewExpenseReporter())

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
