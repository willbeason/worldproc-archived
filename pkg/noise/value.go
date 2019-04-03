package noise

import (
	"math/rand"

	"willbeason/worldproc/pkg/fixed"
)

// Fill generates the underlying noise which Value will interpolate.
//
// src is the source of randomness to use to generate noise.
func (v *Value) Fill(src rand.Source) {
	for i := 0; i < noiseSize2; i++ {
		v.noise[i] = fixed.F16(src.Int63()).Remainder() // 0.0 to 1.0 - 2^-16
	}
}

// Value implements linearly interpolated value noise.
type Value struct {
	// noise is an array of the noise.
	noise [noiseSize2]fixed.F16
}

// At implements Noise.
//
// Guarantees monotonic behavior between integral values.
// Guarantees behavior at (x, y) is equivalent to (x mod size, y mod size)
func (v *Value) At(x, y fixed.F16) fixed.F32 {
	xi, xr := x.Split()
	xi = xi & intMask

	yi, yr := y.Split()
	yi = yi & intMask

	yin := yi << shift
	yip1n := inc(yi) << shift

	// Linearly interpolate based on the four corners of the enclosing square.
	nbl := v.noise[yin+xi]
	nbr := v.noise[yin+inc(xi)]
	nul := v.noise[yip1n+xi]
	nur := v.noise[yip1n+inc(xi)]

	xryr := xr.Times(yr).F16()
	return xryr.Times(nur) +
		(yr - xryr).Times(nul) +
		(xr - xryr).Times(nbr) +
		(fixed.One + xryr - xr - yr).Times(nbl)
}

// inc is a convenient shorthand for adding 1 and taking the modulus by the
// size of the noise to keep the current position in bounds.
func inc(x int) int {
	return (x + 1) & intMask
}
