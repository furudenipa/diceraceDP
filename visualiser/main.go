package main

import (
	"github.com/gdamore/tcell/v2"
)

func main() {

	filepath := "../data/policy.bin"

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
		strides:      computeStrides(),
	}
	s := app.screen
	defer s.Fini()

	app.loading(filepath)

	var debug = false
	if debug {
		app.showOffset()
	}
	app.show()

	for {
		ev := s.PollEvent()
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
			_, maxY := s.Size()
			app.rowViewRange = maxY - 3
			s.Sync()
		}
	}
}

/*
func homain() {
	filePath := "../data/policy.bin"
	data := *reader(filePath)
	fmt.Println("Loaded array from binary file")

	strides := computeStrides()
	reader := bufio.NewReader(os.Stdin)

	for {
		// ユーザーからの入力を取得
		tIndices, err := getUserIndices(reader)
		if err != nil {
			fmt.Printf("入力エラー: %v\n", err)
			continue
		}

		// 入力されたインデックスを取得
		t1, t2, t3, t4, t5, t6 := tIndices[0], tIndices[1], tIndices[2], tIndices[3], tIndices[4], tIndices[5]

		// インデックスの有効性をチェック
		if !validateIndices(t1, t2, t3, t4, t5, t6) {
			fmt.Println("入力されたインデックスが範囲外です。0 <= t1..t6 < 10 を満たしてください。")
			continue
		}

		// データの表示
		for i := 0; i < numSteps; i += 10 {
			end := i + 10
			if end > numSteps {
				end = numSteps
			}

			slice := make([][]byte, end-i)
			for step := i; step < end; step++ {
				stepData := make([]byte, numSquares)
				for square := 0; square < numSquares; square++ {
					idx := getFlatIndex(step, square, t1, t2, t3, t4, t5, t6, strides)
					stepData[square] = data[idx]
				}
				slice[step-i] = stepData
			}

			// スライスを表示
			fmt.Printf("Steps %d to %d:\n", i, end-1)
			for s, stepData := range slice {
				fmt.Printf("  Step %d: %v\n", i+s, stepData)
			}
		}
	}
}*/
