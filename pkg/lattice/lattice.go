package lattice

import (
	"willbeason/worldproc/pkg/fixed"
	"willbeason/worldproc/pkg/noise"
)

type Lattice struct {
	scale fixed.F16
}

func NewLattice(scale fixed.F16) Lattice {
	return Lattice{scale: scale}
}

var _ noise.Source = Lattice{}

// V returns fixed.One32 if the point is a lattice line, otherwise fixed.Zero32.
func (l Lattice) V(x, y fixed.F16) fixed.F32 {
	if ((x % l.scale) < fixed.One16) || ((y % l.scale) < fixed.One16) {
		return fixed.One32
	}
	return fixed.Zero32
}
