package main

import "image/color"

const (
	HEIGHT = 1000
	WIDTH  = 1000
)

var (
	COLOR_BLACK = color.RGBA{}
	COLOR_RED   = color.RGBA{
		R: 255, A: 255,
	}
	COLOR_BLUE = color.RGBA{
		B: 255, A: 255,
	}
	COLOR_GREEN = color.RGBA{
		G: 255, A: 255,
	}
	COLOR_WHITE = color.RGBA{
		R: 255, G: 255, B: 255, A: 255,
	}
)
