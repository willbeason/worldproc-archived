package colors

import (
	"image/color"
)

func Blend(color0, color1 color.NRGBA, p float64) color.NRGBA {
	return color.NRGBA{
		R: uint8((1.0-p)*float64(color0.R) + p*float64(color1.R)),
		G: uint8((1.0-p)*float64(color0.G) + p*float64(color1.G)),
		B: uint8((1.0-p)*float64(color0.B) + p*float64(color1.B)),
		A: uint8((1.0-p)*float64(color0.A) + p*float64(color1.A)),
	}
}
