package main

import (
	"image/color"
	"log"

	"github.com/tfriedel6/canvas"
)

type Grid struct {
	cv     *canvas.Canvas
	R      int // Number of rows in grid
	C      int // Number of columns in grid
	Colors [][]color.RGBA
	h      int // Height of each element
	w      int // Width of each element
}

func NewGrid(cv *canvas.Canvas, R int, C int, clr color.RGBA) *Grid {
	g := Grid{
		cv: cv,
		R:  R,
		C:  C,
		h:  cv.Height() / R,
		w:  cv.Width() / C,
	}
	g.Colors = make([][]color.RGBA, g.R)
	for row := range g.Colors {
		g.Colors[row] = make([]color.RGBA, g.C)
		for column := 0; column < g.C; column += 1 {
			g.Colors[row][column] = clr
		}
	}
	return &g
}

func (g *Grid) isValidCoord(r int, c int) bool {
	if r >= 0 && r < g.R && c >= 0 && c < g.C {
		return true
	}
	return false
}

// Takes in the grid element coordinates and returns the coordinates of top left and width and height
func (g *Grid) toPixels(r int, c int) (float64, float64, float64, float64) {
	if !g.isValidCoord(r, c) {
		log.Printf("Invalid coordinate: %d %d\n", r, c)
		return 0, 0, 0, 0
	}
	x, y := g.w*c, g.h*r
	return float64(x), float64(y), float64(g.w), float64(g.h)
}

func (g *Grid) Render() {
	for r := 0; r < g.R; r += 1 {
		for c := 0; c < g.C; c += 1 {
			clr := g.Colors[r][c]

			g.cv.SetFillStyle(clr.R, clr.G, clr.B, clr.A)
			g.cv.FillRect(g.toPixels(r, c))
		}
	}
}

func (g *Grid) SetColor(r, c int, clr color.RGBA) {
	if g.isValidCoord(r, c) {
		g.Colors[r][c] = clr
	}
}
func (g *Grid) GetColor(r, c int) color.RGBA {
	if g.isValidCoord(r, c) {
		return g.Colors[r][c]
	}
	return COLOR_BLACK
}
