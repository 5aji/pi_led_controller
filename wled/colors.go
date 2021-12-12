package wled

import "math"

// provides utilities to handle color conversion and and creation.

type RGB struct{ R, G, B float64 }

type HSL struct{ H, S, L float64 }

// Verifies if the RGB values are between 0 and 1
func (color *RGB) isValid() bool {

	return color.R >= 0 && color.R <= 1 &&
		color.G >= 0 && color.G <= 1 &&
		color.B >= 0 && color.B <= 1

}

func (color *RGB) Normalize() {
	
	Cmax := math.Max(color.R, math.Max(color.B, color.G)) 
	color.R = color.R/Cmax
	color.G = color.G/Cmax
	color.B = color.B/Cmax
}

func (color *RGB) WledColor() WledColor {
	var b WledColor
	b[0] = byte(math.Round(color.R * 255))
	b[1] = byte(math.Round(color.G * 255))
	b[2] = byte(math.Round(color.B * 255))

	return b

}

func (color *HSL) isValid() bool {

	return color.H >= 0 && color.H <= 1 &&
		color.S >= 0 && color.S <= 1 &&
		color.L >= 0 && color.L <= 1
}
// converts an RGB color to HSL color
func (color *RGB) HSL() HSL {
	Cmax := math.Max(color.R, math.Max(color.B, color.G)) 
	Cmin := math.Min(color.R, math.Min(color.B, color.G)) 

	delta := Cmax - Cmin

	var hue float64

	switch Cmax {
	case color.R:
		hue = ((color.G - color.B)/delta)
	case color.G:
		hue = (2 + (color.B - color.R)/delta)
	case color.B:
		hue = (4 + (color.R - color.G)/delta)
	default:
		hue = 0
	}
	hue = hue / 6
	lightness := (Cmax + Cmin)/2

	var saturation float64

	switch lightness {
	case 0,1:
		saturation = 0
	default:
		saturation = delta/(1 - math.Abs(2 * Cmax - delta - 1))
	}


	return HSL{hue, saturation, lightness}

}



