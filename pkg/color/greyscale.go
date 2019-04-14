package color

import (
	"image/color"
)

func Grayscale() color.Palette {
	result := make([]color.Color, 256)
	result[0] = color.Transparent
	for i := 1; i < 256; i++ {
		result[i] = color.Gray{Y: uint8(i)}
	}
	return color.Palette(result)
}
