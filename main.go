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
	g := NewGrid(cv, 10, 10, COLOR_BLUE)
	wnd.MainLoop(func() {
		g.Render()
	})
}
