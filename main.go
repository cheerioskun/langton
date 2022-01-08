package main

import (
	"log"

	"github.com/tfriedel6/canvas/sdlcanvas"
)

var t int

func main() {
	wnd, cv, err := sdlcanvas.CreateWindow(1000, 1000, "Trial")
	if err != nil {
		panic(err)
	}
	defer wnd.Destroy()
	log.Printf("Canvas dimensions: %d %d", cv.Width(), cv.Height())
	g := NewGrid(cv, 100, 100, COLOR_GREY, 0.0002)
	a := NewAntWithRules(g, 50, 50, RulesLLRR)
	t = 0
	log.Printf("New ant: %v", a)
	stepSize := 100
	wnd.MainLoop(func() {
		for i := 0; i < stepSize; i += 1 {
			t += 1
			a.Move()
		}
		g.Render()
		// time.Sleep(time.Millisecond * 20)
	})
}
