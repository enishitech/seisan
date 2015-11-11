package reporter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tealeg/xlsx"

	"github.com/enishitech/seisan/config"
	"github.com/enishitech/seisan/request"
)

type Reporter interface {
	Report(*xlsx.Sheet, *config.Config, []request.Request) error
}

type SeisanReporter struct {
	reporters []Reporter
}

func New(reporters ...Reporter) *SeisanReporter {
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

func (sr SeisanReporter) Report(baseDir string, target string) error {
	conf, err := config.Load(filepath.Join(baseDir, "config.yaml"))
	if err != nil {
		return err
	}

	reqs, err := request.LoadDir(filepath.Join(baseDir, "data", target))
	if err != nil {
		return err
	}

	targetName := strings.Replace(target, "/", "-", -1)
	xlsx.SetDefaultFont(11, "ＭＳ Ｐゴシック")

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("精算シート")
	if err != nil {
		return err
	}

	renderReportHeader(sheet, targetName, conf.Organization.Name)

	for _, r := range sr.reporters {
		err := r.Report(sheet, conf, reqs)
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
