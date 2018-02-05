package main

import (
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
	table.Rows = getRows()
	table.X = 0
	table.Y = 0
	table.Width = 1
	table.Height = 36
	table.BorderLabel = "CN12306"

	par := ui.NewPar("退出:Q		刷新:R	上:K		下:J")
	par.Height = 3
	par.Width = 17
	par.X = 20
	par.BorderLabel = "帮助"

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, table),
		),
		ui.NewRow(
			ui.NewCol(12, 0, par),
		),
	)

	// Initial Display
	ui.Body.Align()
	ui.Clear()
	ui.Render(ui.Body)
}

func registerEvent() {
	// exit, refresh (force trigger query), down/up (store current view list), timer, resize
	// stop auto-refresh in 10s if you scroll to view the list
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/j", func(ui.Event) {

	})

	ui.Handle("/sys/kbd/k", func(ui.Event) {
	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		table.Rows = getRows()
		ui.Render(ui.Body)
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})
}
