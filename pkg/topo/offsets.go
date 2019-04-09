package topo

import (
	"math/rand"

	"willbeason/proctexture/pkg/noise"

	"willbeason/worldproc/pkg/fixed"
)

// RandomOffsets generates a random set of Offsets from the passed rand.Source.
// Generated offsets are uniformly distributed from 0 to noise.NoiseSize.
func RandomOffsets(src rand.Source) Offsets {
	result := Offsets{}
	for depth := 0; depth < maxDepth; depth++ {
		result.x[depth] = fixed.F16(src.Int63()).Remainder() * noise.NoiseSize
		result.y[depth] = fixed.F16(src.Int63()).Remainder() * noise.NoiseSize
	}
	return result
}

// Offsets is a set of x and y offsets from the origin.
type Offsets struct {
	x [maxDepth]fixed.F16
	y [maxDepth]fixed.F16
}
