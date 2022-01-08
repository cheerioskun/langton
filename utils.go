package main

import (
	"errors"
	"image/color"
	"math"
)

var errInvalidFormat = errors.New("invalid format")

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}

type HSV struct {
	H float64
	S float64
	V float64
}

func RGBtoHSV(c color.RGBA) HSV {
	R, G, B := c.R, c.G, c.B
	var H, S, V float64
	r, g, b := float64(R)/255.0, float64(G)/255.0, float64(B)/255.0
	min := math.Min(math.Min(r, g), b)
	max := math.Max(math.Max(r, g), b)
	del_max := max - min
	V = max
	if del_max == 0 {
		H = 0
		S = 0
	} else {
		S = del_max / max
		del_r := ((max-r)/6 + del_max/2) / del_max
		del_g := ((max-g)/6 + del_max/2) / del_max
		del_b := ((max-b)/6 + del_max/2) / del_max
		if r == max {
			H = del_b - del_g
		} else if g == max {
			H = (del_r - del_b) / 3
		} else if b == max {
			H = (2 * (del_g - del_b)) / 3
		}
		if H < 0 {
			H += 1
		} else if H > 1 {
			H -= 1
		}
	}
	return HSV{
		H: H,
		S: S,
		V: V,
	}
}

func HSVtoRGB(hsv HSV) color.RGBA {
	H, S, V := hsv.H, hsv.S, hsv.V
	var R, G, B uint8
	if S == 0 {
		R = uint8(V * 255)
		G = uint8(V * 255)
		B = uint8(V * 255)
	} else {
		h := H * 6
		if h == 6 {
			h = 0
		}
		var_i := math.Floor(h)
		var_1 := V * (1 - S)
		var_2 := V * (1 - S*(h-var_i))
		var_3 := V * (1 - S*(1-(h-var_i)))
		var var_r, var_g, var_b float64
		if var_i == 0 {
			var_r = V
			var_g = var_3
			var_b = var_1
		} else if var_i == 1 {
			var_r = var_2
			var_g = V
			var_b = var_1
		} else if var_i == 2 {
			var_r = var_1
			var_g = V
			var_b = var_3
		} else if var_i == 3 {
			var_r = var_1
			var_g = var_2
			var_b = V
		} else if var_i == 4 {
			var_r = var_3
			var_g = var_1
			var_b = V
		} else {
			var_r = V
			var_g = var_1
			var_b = var_2
		}

		R = uint8(var_r * 255)
		G = uint8(var_g * 255)
		B = uint8(var_b * 255)
	}
	return color.RGBA{
		R: R,
		G: G,
		B: B,
		A: 255,
	}
}

func ExpInterpolate(c1, c2 HSV, t float64) HSV {
	res := HSV{}
	res.H = c1.H*(1-t) + c2.H*t
	res.S = c1.S*(1-t) + c2.S*t
	res.V = c1.V*(1-t) + c2.V*t
	return res
}
