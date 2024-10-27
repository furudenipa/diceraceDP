package visualizer

import (
	"log/slog"

	"github.com/furudenipa/diceraceDP/pkg/reader"
	"github.com/gdamore/tcell/v2"
)

func Run(filepath string) {

	screen, err := tcell.NewScreen()
	if err != nil {
		slog.Error("Failed to create new screen", slog.String("error", err.Error()))
		return
	}
	if err := screen.Init(); err != nil {
		slog.Error("Failed to initialize screen", slog.String("error", err.Error()))
		return

	}
	_, maxY := screen.Size()
	app := &App{
		screen:           screen,
		filepath:         filepath,
		policy:           make([]byte, 0),
		rowIndex:         0,
		rowViewRange:     maxY - appOffsetY,
		remainingTickets: make([]int, 6),
		strides:          reader.ComputeStrides(),
	}
	defer screen.Fini()

	app.loading()

	app.drawFilepath()
	app.render()

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
				app.render()

			case tcell.KeyEscape, tcell.KeyCtrlC:
				// EscapeキーかCtrl-Cでループを終了
				return

			case tcell.KeyDown:
				app.incrementRowIndex()
				app.render()

			case tcell.KeyUp:
				app.decrementRowIndex()
				app.render()
			}

		case *tcell.EventResize:
			screen.Sync()
			_, maxY := screen.Size()
			app.rowViewRange = maxY - appOffsetY
			app.render()
		}
	}
}
