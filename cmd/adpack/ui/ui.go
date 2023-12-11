package main

import (
	"context"

	"github.com/nsf/termbox-go"
)

func UI(ctx context.Context) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	x, y := 0, 0

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x, y, 'X', termbox.ColorDefault, termbox.ColorDefault)
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainLoop
			case termbox.KeyArrowUp, termbox.Key('k'):
				if y > 0 {
					y--
				}
			case termbox.KeyArrowDown, termbox.Key('j'):
				if _, h := termbox.Size(); y < h {
					y++
				}
			case termbox.KeyArrowLeft, termbox.Key('h'):
				if x > 0 {
					x--
				}
			case termbox.KeyArrowRight, termbox.Key('l'):
				if w, _ := termbox.Size(); x < w {
					x++
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
