package main

import (
	"log"

	"github.com/tfriedel6/canvas/sdlcanvas"
)

func main() {
	wnd, cv, err := sdlcanvas.CreateWindow(1000, 1000, "Trial")
	if err != nil {
		panic(err)
	}
	defer wnd.Destroy()
	log.Printf("Canvas dimensions: %d %d", cv.Width(), cv.Height())
	g := NewGrid(cv, 100, 100, COLOR_WHITE)
	a := NewAntWithRules(g, 50, 50, RulesLLRR)
	log.Printf("New ant: %v", a)
	wnd.MainLoop(func() {
		a.Move()
		g.Render()
		// time.Sleep(time.Millisecond * 20)
	})
}
