package lattice

import (
	"willbeason/worldproc/pkg/fixed"
)

type Lattice struct {
	invScale fixed.F16
}

func NewLattice(scale fixed.F16) Lattice {
	return Lattice{
		invScale: scale.Invert(),
	}
}

// Dist returns the distance from the position to the nearest lattice edge.
func (l Lattice) Dist(x, y int) fixed.F16 {
	xr, yr := l.invScale.TimesInt(x).Remainder(), l.invScale.TimesInt(y).Remainder()
	if xr > fixed.Half16 {
		xr = fixed.One16 - xr
	}
	if yr > fixed.Half16 {
		yr = fixed.One16 - yr
	}
	if xr < yr {
		return xr
	}
	return yr
}
