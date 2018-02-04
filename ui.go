package main

import (
	"fmt"

	ui "github.com/gizak/termui"
)

// terminal UI
func startUI() {
	results := q(queryURL)
	// fmt.Println("length", len(results))

	rows := parseResult(results)

	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	table := ui.NewTable()
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
	fmt.Println("finished!")
	// calculate layout
	ui.Body.Align()

	ui.Clear()
	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/timer/5s", func(e ui.Event) {
		// t := e.Data.(ui.EvtTimer)
		// i := t.Count
		// if i > 103 {
		// 	ui.StopLoop()
		// 	return
		// }

		ui.Clear()
		ui.Body.Rows = nil
		ui.Render(ui.Body)
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Loop()
}
