package main_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/franela/goblin"
	"github.com/gdamore/tcell/termbox"
)

// go test -v -count=1 -timeout 30s -run ^Test$ gounchpad

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("main test", func() {

		g.It("aaa", func() {

			err := termbox.Init()
			if err != nil {
				g.Assert(err)
			}

			defer termbox.Close()

			for {
				switch ev := termbox.PollEvent(); ev.Type {
				case termbox.EventKey:
					switch ev.Key {
					case termbox.KeyEsc:
						fmt.Println("end")
						return
					case termbox.KeyArrowUp:
						fmt.Println("Up")
					case termbox.KeyArrowDown:
						fmt.Println("Down")
					case termbox.KeyArrowLeft:
						fmt.Println("Left")
					case termbox.KeyArrowRight:
						fmt.Println("Right")
					case termbox.KeySpace:
						fmt.Println("Space")
					default:
						fmt.Println("other key:", ev.Key)
					}
				default:
					log.Println("other")
				}
			}

		})

	})
}
