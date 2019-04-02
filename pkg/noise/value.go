package noise

import (
	"fmt"
	"math/rand"
	"time"

	"willbeason/worldproc/pkg/fixed"
)

// NewValue returns a Noise whose underlying implementation implements interpolated value noise.
// Repeats infinitely and regularly by taking the modulus of the passed position.
// Non-integral values are linearly interpolated.
//
// src is the source of randomness to use to generate noise. If nil, creates a seed from the
// current timestamp.
func NewValue(src rand.Source) *Value {
	if src == nil {
		seed := time.Now().UnixNano()
		fmt.Println("Noise seed: ", seed)
		src = rand.NewSource(seed)
	}

	result := &Value{}
	for i := 0; i < noiseSize2; i++ {
		// TODO: test whether bitwise and is more efficient than bit shifting.
		result.noise[i] = fixed.F16(src.Int63()).Remainder() // 0.0 to 1.0 - 2^-16
	}
	return result
}

// Value implements linearly interpolated value noise.
type Value struct {
	// noise is an array of the noise.
	noise [noiseSize2]fixed.F16
}

// At implements Noise.
func (n *Value) At(x, y fixed.F16) fixed.F32 {
	xi, xr := x.Split()
	xi = xi & intMask

	yi, yr := y.Split()
	yi = yi & intMask

	yin := yi << shift
	yip1n := inc(yi) << shift

	// Linearly interpolate based on the four corners of the enclosing square.
	nbl := n.noise[yin+xi]
	nbr := n.noise[yin+inc(xi)]
	nul := n.noise[yip1n+xi]
	nur := n.noise[yip1n+inc(xi)]

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
