package main

import (
	"errors"
	"log"

	ui "github.com/gizak/termui"
)

const (
	MaxRecordSize = 1000
)

var (
	win = Window{}
)

// ScrollTable is a scrollable table
type ScrollTable struct {
	ui.Table
	Head         []string
	Body         [][]string
	currentLine  int
	currentIndex int
	RowPerPage   int
}

func (t *ScrollTable) moveUp() {

}

func (t *ScrollTable) moveDown() {

}

func (t *ScrollTable) prevPage() {

}

func (t *ScrollTable) nextPage() {

}

func (t *ScrollTable) update(rows [][]string) error {
	if rows == nil || len(rows) == 0 {
		if t.Body == nil {
			t.Rows = [][]string{t.Head}
		}
		return errors.New("updating rows with empty data")
	}

	if len(rows) > MaxRecordSize {
		t.Body = rows[:MaxRecordSize]
	} else {
		t.Body = rows
	}
	t.currentIndex = 0
	t.currentLine = 0

	rs := [][]string{}
	if t.Head != nil && len(t.Head) > 0 {
		rs = append(rs, t.Head)
	}

	currentPageSize := 12
	if len(t.Body) < t.RowPerPage {
		currentPageSize = len(t.Body)
	}
	rs = append(rs, t.Body[:currentPageSize]...)

	t.Rows = rs
	return nil
}

// HintPar stores the hints and status message
type HintPar struct {
	ui.Par
	Message string
}

// Window is the structure stores the whole ui widgets
type Window struct {
	TicketTable ScrollTable
	HintBar     HintPar
	inited      bool
}

func (w *Window) Init() {
	if w.inited {
		return
	}

	w.renderWidgets()
	w.registerEvents()
	w.inited = true
}

func (w *Window) Refresh() {
	if w.inited {
		ui.Render(ui.Body)
	}
}

func (w *Window) renderWidgets() {
	w.TicketTable = ScrollTable{
		Table:      *ui.NewTable(),
		Head:       []string{"车次", "起始", "--->", "终止", "发车时间", "--", "到达时间", "历时", "座位信息", "", "", ""},
		RowPerPage: 12,
	}
	w.TicketTable.X = 0
	w.TicketTable.Y = 0
	w.TicketTable.Width = 1
	w.TicketTable.Height = 36
	w.TicketTable.BorderLabel = "CN12306"
	err := w.TicketTable.update(getRows())
	if err != nil {
		log.Println(err)
	}

	w.HintBar = HintPar{Par: *ui.NewPar("退出:Q		刷新:R	上:K		下:J")}
	w.HintBar.Height = 3
	w.HintBar.Width = 17
	w.HintBar.X = 20
	w.HintBar.BorderLabel = "帮助"

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, &w.TicketTable.Table)),
		ui.NewRow(ui.NewCol(12, 0, &w.HintBar.Par)),
	)

	ui.Body.Align()
	ui.Clear()
	ui.Render(ui.Body)
}

// exit, refresh (force trigger query), down/up (store current view list), timer, resize
// stop auto-refresh in 10s if you scroll to view the list
func (w *Window) registerEvents() {
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/r", func(ui.Event) {
		w.TicketTable.update(getRows())
		w.Refresh()
	})

	// prev page
	ui.Handle("/sys/kbd/h", func(ui.Event) {

	})

	// scroll down
	ui.Handle("/sys/kbd/j", func(ui.Event) {
		// if index+12 < len(rows) {
		// 	index++
		// }
		// end := index + 12
		// if len(rows) < 12 {
		// 	end = len(rows) - 1
		// }
		// w.TicketTable.Rows = rows[index:end]
		// ui.Render(ui.Body)
	})

	// scroll up
	ui.Handle("/sys/kbd/k", func(ui.Event) {
		// if index > 1 {
		// 	index--
		// }
		// end := index + 12
		// if len(rows) < 12 {
		// 	end = len(rows) - 1
		// }
		// w.TicketTable.Rows = rows[index:end]
		// ui.Render(ui.Body)
	})

	// next page page
	ui.Handle("/sys/kbd/l", func(ui.Event) {

	})

	// gg jumps to start
	ui.Handle("/sys/kbd/g", func(e ui.Event) {

	})

	// GG jumps to end
	ui.Handle("/sys/kbd/G", func(e ui.Event) {

	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		// table.Rows = getRows()
		// ui.Render(ui.Body)
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})
}

// terminal UI
func startUI() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	win.Init()
	ui.Loop()
}
