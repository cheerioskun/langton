package main

import (
	"log"
	"math"

	"github.com/tfriedel6/canvas"
)

type Grid struct {
	cv        *canvas.Canvas
	R         int // Number of rows in grid
	C         int // Number of columns in grid
	Colors    [][]Color
	ColorTime [][]int
	Bg        Color
	h         int     // Height of each element
	w         int     // Width of each element
	decay     float64 // Decay factor for the colors
}

func NewGrid(cv *canvas.Canvas, R int, C int, clr Color, decay float64) *Grid {
	g := Grid{
		cv:    cv,
		R:     R,
		C:     C,
		h:     cv.Height() / R,
		w:     cv.Width() / C,
		Bg:    clr,
		decay: decay,
	}
	g.Colors = make([][]Color, g.R)
	g.ColorTime = make([][]int, g.R)
	for row := range g.Colors {
		g.Colors[row] = make([]Color, g.C)
		g.ColorTime[row] = make([]int, g.C)
		for column := 0; column < g.C; column += 1 {
			g.Colors[row][column] = clr
			g.ColorTime[row][column] = 0
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
			crgba, _ := ParseHexColor(string(clr))
			bg, _ := ParseHexColor(string(g.Bg))
			chsv, bhsv := RGBtoHSV(crgba), RGBtoHSV(bg)
			ratio := 1 - math.Exp(-float64(t-g.ColorTime[r][c])*g.decay)
			chsv = ExpInterpolate(chsv, bhsv, ratio)
			crgba = HSVtoRGB(chsv)
			g.cv.SetFillStyle(crgba.R, crgba.G, crgba.B, crgba.A)
			g.cv.FillRect(g.toPixels(r, c))
		}
	}
}

func (g *Grid) SetColor(r, c int, clr Color) {
	if g.isValidCoord(r, c) {
		g.Colors[r][c] = clr
		g.ColorTime[r][c] = t
	}
}
func (g *Grid) GetColor(r, c int) Color {
	if g.isValidCoord(r, c) {
		return g.Colors[r][c]
	}
	return COLOR_BLACK
}
