package colors

import (
	"image/color"
	"math"
)

type Scale struct {
	colors []color.NRGBA
}

func (s *Scale) At(f float64) color.NRGBA {
	idx := int(math.Floor(f))

	if idx < 0 {
		return s.colors[0]
	} else if idx+1 >= len(s.colors) {
		return s.colors[len(s.colors)-1]
	}

	r := f - math.Floor(f)
	return Blend(s.colors[idx], s.colors[idx+1], r)
}
