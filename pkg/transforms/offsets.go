package transforms

import (
	"math/rand"

	"willbeason/proctexture/pkg/noise"

	"willbeason/worldproc/pkg/fixed"
)

type offset struct {
	X, Y fixed.F16
}

var (
	NoOffset = offset{X: 0, Y: 0}
)

// Offsets is a set of x and y offsets from the origin.
type Offsets [maxDepth]offset

// RandomOffsets generates a random set of Offsets from the passed rand.Source.
// Generated offsets are uniformly distributed from 0 to noise.NoiseSize.
func RandomOffsets(src rand.Source) Offsets {
	result := Offsets{}
	for depth := 0; depth < maxDepth; depth++ {
		result[depth].X = fixed.F16(src.Int63()).Remainder() * noise.NoiseSize
		result[depth].Y = fixed.F16(src.Int63()).Remainder() * noise.NoiseSize
	}
	return result
}

func Offset(x, y fixed.F16) offset {
	return offset{X: x, Y: y}
}
