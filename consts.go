package main

const (
	HEIGHT = 1000
	WIDTH  = 1000
)

type Color string

var (
	COLOR_BLACK     Color = "#000000"
	COLOR_RED       Color = "#FF0000"
	COLOR_BLUE      Color = "#0000FF"
	COLOR_GREEN     Color = "#00FF00"
	COLOR_WHITE     Color = "#FFFFFF"
	COLOR_GREY      Color = "#272124"
	COLOR_LIME      Color = "#B9F621"
	COLOR_TURQUOISE Color = "#10ADC5"
	COLOR_OLIVE     Color = "#3D6341"
)

var (
	RulesLLRR = map[Color]Move{
		COLOR_WHITE: {
			Rotation: LeftRotation,
			Recolor:  COLOR_LIME,
		},
		COLOR_LIME: {
			Rotation: LeftRotation,
			Recolor:  COLOR_TURQUOISE,
		},
		COLOR_TURQUOISE: {
			Rotation: RightRotation,
			Recolor:  COLOR_GREY,
		},
		COLOR_GREY: {
			Rotation: RightRotation,
			Recolor:  COLOR_WHITE,
		},
	}
)
