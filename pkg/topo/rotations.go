package topo

import (
	"math"
	"math/rand"

	"willbeason/worldproc/pkg/fixed"
)

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

type Rotations struct {
	sin [maxDepth]fixed.F16
	cos [maxDepth]fixed.F16
}
