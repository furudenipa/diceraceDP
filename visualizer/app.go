package visualizer

import (
	"fmt"
	"strconv"

	"github.com/furudenipa/diceraceDP/config"
	"github.com/furudenipa/diceraceDP/reader"
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

const (
	appOffsetX = 6
	appOffsetY = 4
)

var Colors = []tcell.Color{
	tcell.GetColor("#D0D0D0"), // roll
	tcell.GetColor("#E14210"), // 1
	tcell.GetColor("#E78A11"), // 2
	tcell.GetColor("#70CF6F"), // 3
	tcell.GetColor("#295EBE"), // 4
	tcell.GetColor("#5A328E"), // 5
	tcell.GetColor("#C33AA8"), // 6
	tcell.ColorBlack,          // None
}

type App struct {
	screen       tcell.Screen
	policy       []byte
	rowIndex     int
	rowViewRange int
	ticketIndexs []int
	strides      []int
}

// loading画面の表示とデータの読み込み
func (app *App) loading(filepath string) {
	s := app.screen
	s.Clear()

	// loading画面の表示
	app.SetContents(appOffsetX, appOffsetY, "Loading...", tcell.StyleDefault)
	s.Show()

	// readerを使ってデータを読む
	app.policy = *reader.PolicyReader(filepath)
}

// 画面表示、クリア不要であること
func (app *App) show() {
	s := app.screen
	app.drawRow()
	app.drawColumn()
	app.drawLogo()
	app.drawMatrix()
	app.drawTickets()
	s.Show()
}

func (app *App) drawMatrix() {

	// 現在のt1〜t6の値を取得
	t1, t2, t3, t4, t5, t6 := app.ticketIndexs[0], app.ticketIndexs[1], app.ticketIndexs[2], app.ticketIndexs[3], app.ticketIndexs[4], app.ticketIndexs[5]

	for step := 0; step < app.rowViewRange; step++ {
		for square := 0; square < config.NumSquares; square++ {
			var color tcell.Color
			if step+app.rowIndex < config.NumSteps {
				idx := reader.GetFlatIndex(step+app.rowIndex, square, t1, t2, t3, t4, t5, t6, app.strides)
				value := int(app.policy[idx] % 8)
				color = Colors[value]
			} else {
				color = Colors[7] //policyが計算されていない場合
			}

			/*if square == 12 {
				color = colors[7]
			}*/

			// スクリーン上の位置を計算
			x := appOffsetX + square*3
			y := appOffsetY + step

			// 四角形を描画
			app.drawSquare(x, y, color)
		}
	}
}

func (app *App) drawRow() {
	s := app.screen
	for i := 0; i < app.rowViewRange; i++ {
		// 現在表示する行のインデックスを取得
		rowIndex := app.rowIndex + i
		// インデックスを文字列に変換, 3桁になるように左にスペースを追加
		indexStr := fmt.Sprintf("%3d", rowIndex)

		// 数値を右揃えにするための開始x座標を計算
		startX := -(len(indexStr) + 1)

		for j, char := range indexStr {
			x := startX + j + appOffsetX
			y := i + appOffsetY
			s.SetContent(x, y, char, nil, tcell.StyleDefault)
		}
	}
}

func (app *App) drawColumn() {
	s := app.screen
	y := appOffsetY - 1

	for columnIndex := 0; columnIndex < config.NumSquares; columnIndex++ {
		indexStr := fmt.Sprintf("%d", columnIndex)
		startX := -(len(indexStr) + 1)
		for j, char := range indexStr {
			x := appOffsetX + (columnIndex+1)*3 + startX + j // 各桁のx座標を計算
			s.SetContent(x, y, char, nil, tcell.StyleDefault)
		}

	}
}

func (app *App) drawTickets() {

	s := app.screen
	offset := appOffsetX + 9
	x := offset
	for ticket := 0; ticket < 6; ticket++ {
		ticketStr := fmt.Sprintf("%d", app.ticketIndexs[ticket])
		app.SetContents(x, 1, " T"+strconv.Itoa(ticket+1)+":", tcell.StyleDefault)
		x += 4
		for _, char := range ticketStr {
			s.SetContent(x, 1, char, nil, tcell.StyleDefault.Background(Colors[ticket+1]).Foreground(tcell.ColorBlack))
			x++
		}
	}
}

func (app *App) drawLogo() {
	app.SetContents(2, 0, "Dice", tcell.StyleDefault.Foreground(tcell.ColorBlue))
	app.SetContents(6, 0, "Race", tcell.StyleDefault.Foreground(tcell.ColorWhite))
	app.SetContents(2, 1, "Visualiser", tcell.StyleDefault.Foreground(tcell.GetColor("#D0D0D0")))
}

func (app *App) showOffset() {
	// offsetX - offsetY
	_, maxY := app.screen.Size()
	for x := 0; x < appOffsetX; x++ {
		for y := 0; y < appOffsetY; y++ {
			app.screen.SetContent(x, y, '/', nil, tcell.StyleDefault.Background(tcell.ColorGray))
		}
	}

	// other - offsetY
	for x := appOffsetX; x < 3*config.NumSquares+appOffsetX; x++ {
		for y := 0; y < appOffsetY; y++ {
			app.screen.SetContent(x, y, '/', nil, tcell.StyleDefault.Background(tcell.ColorRed))
		}
	}

	// offsetX - other
	for x := 0; x < appOffsetX; x++ {
		for y := appOffsetY; y < maxY; y++ {
			app.screen.SetContent(x, y, '/', nil, tcell.StyleDefault.Background(tcell.ColorGreen))
		}
	}
}

func (app *App) drawSquare(x, y int, c tcell.Color) {
	s := app.screen
	s.SetContent(x, y, '▌', nil, tcell.StyleDefault.Background(c).Foreground(tcell.ColorBlack))
	s.SetContent(x+1, y, ' ', nil, tcell.StyleDefault.Background(c).Foreground(tcell.ColorBlack))
	s.SetContent(x+2, y, '▐', nil, tcell.StyleDefault.Background(c).Foreground(tcell.ColorBlack))
}

func (app *App) incrementTicket(index int) {
	app.ticketIndexs[index]++
	if app.ticketIndexs[index] >= config.MaxTickets {
		app.ticketIndexs[index] = 0
	}
}

func (app *App) decrementTicket(index int) {
	app.ticketIndexs[index]--
	if app.ticketIndexs[index] < 0 {
		app.ticketIndexs[index] = config.MaxTickets - 1
	}
}

func (app *App) incrementRowIndex() {
	app.rowIndex++
	if app.rowIndex >= config.NumSteps {
		app.rowIndex = config.NumSteps
	}
}

func (app *App) decrementRowIndex() {
	app.rowIndex--
	if app.rowIndex < 0 {
		app.rowIndex = 0
	}
}

func (app *App) SetContents(x, y int, text string, style tcell.Style) {
	s := app.screen
	for _, r := range text {
		s.SetContent(x, y, r, nil, style)
		x += runewidth.RuneWidth(r)
	}
}
