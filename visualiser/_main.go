package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func main() {
	fmt.Println("Settings")
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

}
