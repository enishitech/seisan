package main

import (
	"fmt"
	"sort"

	"github.com/tealeg/xlsx"
)

type ExpenseReport struct {
	summary map[string]int
	lines   []Expense
}

func newExpenseReport(seisanRequests []SeisanRequest) *ExpenseReport {
	report := &ExpenseReport{}
	report.makeSummary(seisanRequests)
	report.makeLines(seisanRequests)
	return report
}

func (self *ExpenseReport) makeSummary(seisanRequests []SeisanRequest) {
	self.summary = map[string]int{}
	for _, seisanRequest := range seisanRequests {
		for _, expense := range seisanRequest.Expenses {
			if _, exists := self.summary[seisanRequest.Applicant]; exists {
				self.summary[seisanRequest.Applicant] = expense.Amount
			} else {
				self.summary[seisanRequest.Applicant] += expense.Amount
			}
		}
	}
}

type ByDate []Expense

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date < a[j].Date }

func (self *ExpenseReport) makeLines(in []SeisanRequest) {
	self.lines = []Expense{}
	for _, seisanRequest := range in {
		for _, expense := range seisanRequest.Expenses {
			expense.Applicant = seisanRequest.Applicant
			self.lines = append(self.lines, expense)
		}
	}
	sort.Sort(ByDate(self.lines))
	fmt.Printf("Processed %d expenses\n", len(self.lines))
}

func (self *ExpenseReport) renderSummary(sheet *xlsx.Sheet) {
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
	for key, value := range self.summary {
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

func (self *ExpenseReport) renderLines(sheet *xlsx.Sheet) {
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
	for _, detail := range self.lines {
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

func (self *ExpenseReport) render(sheet *xlsx.Sheet) {
	self.renderSummary(sheet)
	self.renderLines(sheet)
}
