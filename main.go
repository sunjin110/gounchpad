package main

import (
	"fmt"
	"gounchpad/pkg/sound"

	"github.com/gdamore/tcell/termbox"
)

const coldef = termbox.ColorDefault

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			// sound
			go sound.Sound(ev.Ch)

			// // 画面クリア
			termbox.Clear(coldef, coldef)
			termbox.SetCell(10, 10, ev.Ch, termbox.ColorWhite, termbox.AttrBold)
			termbox.Flush()

			// esqで終了
			if ev.Key == termbox.KeyEsc {
				fmt.Println("bye...")
				return
			}
		}
	}

}
