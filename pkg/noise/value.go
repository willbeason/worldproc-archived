package noise

import (
	"math/rand"

	"willbeason/worldproc/pkg/fixed"
)

// Fill generates the underlying noise which Value will interpolate.
//
// src is the source of randomness to use to generate noise.
func (v *Value) Fill(src rand.Source) {
	for i := 0; i < size2; i++ {
		v.noise[i] = fixed.F16(src.Int63()).Remainder() // 0.0 to 1.0 - 2^-16
	}
}

// Value implements linearly interpolated value noise.
type Value struct {
	// noise is an array of the underlying noise.
	// Row (x) is the [0,size) bits of the index.
	// Column (y) is the [size,2*size) bits of the index.
	noise [size2]fixed.F16
}

// V implements Source.
//
// Guarantees monotonic behavior between integral values.
// Guarantees behavior at (x, y) is equivalent to (x mod size, y mod size)
func (v *Value) V(x, y fixed.F16) fixed.F32 {
	// Take the modulus of the integral parts of each coordinate.
	// Each measured faster stored rather than recomputed 4 times.
	xi := x.Int() & intMask
	yi := (int(y) >> revShift) & int2Mask

	// Get the value at each corner surrounding the position.
	// The compiler optimizes away these assignments; this is for readability.
	//
	// The additions and bitwise-anding are offsets and moduli respectively.
	// Measured faster inlined rather than as a function or stored.
	vBottomLeft := v.noise[yi+xi]
	vBottomRight := v.noise[yi+((xi+1)&intMask)]
	vUpperLeft := v.noise[((yi+size)&int2Mask)+xi]
	vUpperRight := v.noise[((yi+size)&int2Mask)+((xi+1)&intMask)]

	// Linearly interpolate based on the four corners of the enclosing square.

	// Measured faster to store these rather than recompute.
	xr := x.Remainder()
	yr := y.Remainder()
	// Measured faster to store xryr as to
	// 1) try to eliminate the second use, or
	// 2) not store the value.
	xryr := xr.Times(yr).F16()
	return xryr.Times(vUpperRight) +
		(yr - xryr).Times(vUpperLeft) +
		(xr - xryr).Times(vBottomRight) +
		(fixed.One + xryr - xr - yr).Times(vBottomLeft)
}
