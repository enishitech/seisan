package expense

import (
	"sort"

	"github.com/tealeg/xlsx"

	"github.com/enishitech/seisan/config"
	"github.com/enishitech/seisan/request"
)

type ExpenseReporter struct{}

func NewReporter() *ExpenseReporter {
	return &ExpenseReporter{}
}

type ExpenseRequest struct {
	Applicant       string  `yaml:"applicant"`
	ExpeneseEntries []Entry `yaml:"expense"`
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

type ByDate []Entry

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date < a[j].Date }

func (reporter ExpenseReporter) renderEntries(sheet *xlsx.Sheet, entries []Entry) {
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
	entries := make([]Entry, 0)
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
