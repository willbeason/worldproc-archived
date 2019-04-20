package topo

import (
	"image"

	"willbeason/worldproc/pkg/fixed"
)

type Topography struct {
	Pix    []fixed.F32
	Stride int
}

// NewTopography instantiates a Topography with all values 0.
func NewTopography(rect image.Rectangle, f func(x, y int) fixed.F32) Topography {
	result := Topography{
		Pix:    make([]fixed.F32, rect.Max.X*rect.Max.Y),
		Stride: rect.Max.X,
	}
	for x := 0; x < rect.Max.X; x++ {
		for y := 0; y < rect.Max.Y; y++ {
			result.Pix[y*rect.Max.X+x] = f(x, y)
		}
	}
	return result
}

func (t *Topography) Height(x, y int) fixed.F32 {
	return t.Pix[y*t.Stride+x]
}
