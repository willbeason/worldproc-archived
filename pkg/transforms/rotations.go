package transforms

import (
	"math"
	"math/rand"

	"willbeason/worldproc/pkg/fixed"
)

type rotation struct {
	Sin, Cos fixed.F16
}

// Rotations is the sine and cosine of rotations about the origin.
type Rotations [maxDepth]rotation

// RandomRotations generates a random set of Rotations from the passed
// rand.Source. Generated rotations are approximately uniformly distributed.
func RandomRotations(src rand.Source) Rotations {
	result := Rotations{}
	r := rand.New(src)
	for depth := 0; depth < maxDepth; depth++ {
		rot := r.Float64() * 2 * math.Pi
		result[depth].Sin = fixed.Float(math.Sin(rot))
		result[depth].Cos = fixed.Float(math.Cos(rot))
	}
	return result
}
