package visualizer

import (
	"github.com/furudenipa/diceraceDP/reader"
	"github.com/gdamore/tcell/v2"
)

func Run(filepath string) {

	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	_, maxY := screen.Size()
	app := &App{
		screen:       screen,
		policy:       make([]byte, 10),
		rowIndex:     0,
		rowViewRange: maxY - 3,
		ticketIndexs: make([]int, 6),
		strides:      reader.ComputeStrides(),
	}
	defer screen.Fini()

	app.loading(filepath)

	var debug = false
	if debug {
		app.showOffset()
	}
	app.show()

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				switch ev.Rune() {
				case '1':
					app.incrementTicket(0)
				case '2':
					app.incrementTicket(1)
				case '3':
					app.incrementTicket(2)
				case '4':
					app.incrementTicket(3)
				case '5':
					app.incrementTicket(4)
				case '6':
					app.incrementTicket(5)

				case '!':
					app.decrementTicket(0)
				case '"':
					app.decrementTicket(1)
				case '#':
					app.decrementTicket(2)
				case '$':
					app.decrementTicket(3)
				case '%':
					app.decrementTicket(4)
				case '&':
					app.decrementTicket(5)
				}
				app.show()

			case tcell.KeyEscape, tcell.KeyCtrlC:
				// EscapeキーかCtrl-Cでループを終了
				return

			case tcell.KeyDown:
				app.incrementRowIndex()
				app.show()

			case tcell.KeyUp:
				app.decrementRowIndex()
				app.show()
			}

		case *tcell.EventResize:
			screen.Sync()
			_, maxY := screen.Size()
			app.rowViewRange = maxY - 3
			app.show()
		}
	}
}
