package main

import (
	"github.com/artificerpi/cn12306/query"
	ui "github.com/gizak/termui"
)

// TODO should be set property
type TicketRow [3]string

var (
	table *ui.Table
)

// terminal UI
func startUI() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	resetUI()
	registerEvent()
	ui.Loop()
}

func resetUI() {
	table = ui.NewTable()
	results := query.Q(queryURL)
	rows := parseResult(results)
	table.Rows = rows
	table.X = 0
	table.Y = 0
	table.Width = 1
	table.Height = 36

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, table),
		),
	)

	// Initial Display
	ui.Body.Align()
	ui.Clear()
	ui.Render(ui.Body)
}

func registerEvent() {
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/timer/1s", func(e ui.Event) {
		results := query.Q(queryURL)
		rows := parseResult(results)
		table.Rows = rows
		ui.Render(ui.Body)
	})

	// TODO add update event

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})
}
