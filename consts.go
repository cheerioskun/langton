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
	COLOR_RED_LIGHT = color.RGBA{
		R: 255, G: 50, B: 50, A: 255,
	}
	COLOR_BLUE_LIGHT = color.RGBA{
		R: 50, G: 50, B: 255, A: 255,
	}
)

var (
	RulesLLRR = map[color.RGBA]Move{
		COLOR_WHITE: {
			Rotation: LeftRotation,
			Recolor:  COLOR_RED_LIGHT,
		},
		COLOR_RED_LIGHT: {
			Rotation: LeftRotation,
			Recolor:  COLOR_RED,
		},
		COLOR_RED: {
			Rotation: RightRotation,
			Recolor:  COLOR_BLUE,
		},
		COLOR_BLUE: {
			Rotation: RightRotation,
			Recolor:  COLOR_WHITE,
		},
	}
)
