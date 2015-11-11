package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

type SeisanReport struct {
	config        Config
	expenseReport *ExpenseReport
}

func newSeisanReport(seisanRequests []SeisanRequest, config Config) *SeisanReport {
	report := &SeisanReport{}
	report.config = config
	report.expenseReport = newExpenseReport(seisanRequests)
	return report
}

func (self *SeisanReport) renderReportHeader(sheet *xlsx.Sheet, name string) {
	var row *xlsx.Row
	var cell *xlsx.Cell

	row = sheet.AddRow()
	cell = row.AddCell()
	orgName := fmt.Sprint(self.config.Organization["name"])
	cell.SetValue(orgName + " 精算シート " + name)
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.SetValue("作成時刻")
	cell = row.AddCell()
	cell.SetValue(time.Now())
	row = sheet.AddRow()
	cell = row.AddCell()
}

func (self *SeisanReport) export() error {
	targetName := strings.Replace(self.config.Target, "/", "-", -1)

	xlsx.SetDefaultFont(11, "ＭＳ Ｐゴシック")

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("精算シート")
	if err != nil {
		return err
	}

	self.renderReportHeader(sheet, targetName)
	self.expenseReport.render(sheet)

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
