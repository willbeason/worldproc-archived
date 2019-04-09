package topo

import (
	"math"
	"math/rand"

	"willbeason/worldproc/pkg/fixed"
)

// RandomRotations generates a random set of Rotations from the passed
// rand.Source. Generated rotations are approximately uniformly distributed.
func RandomRotations(src rand.Source) Rotations {
	result := Rotations{}
	r := rand.New(src)
	for depth := 0; depth < maxDepth; depth++ {
		rot := r.Float64() * 2 * math.Pi
		result.sin[depth] = fixed.Float(math.Sin(rot))
		result.cos[depth] = fixed.Float(math.Cos(rot))
	}
	return result
}

// Rotations is the sine and cosine of rotations about the origin.
type Rotations struct {
	sin [maxDepth]fixed.F16
	cos [maxDepth]fixed.F16
}
