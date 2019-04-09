package topo

import (
	"willbeason/worldproc/pkg/fixed"
)

// PowerScales generates a set of scales where the factor from each scale to its
// neighbors is scale, with maxScale as the first.
func PowerScales(maxScale, scale fixed.F16) Scales {
	result := Scales{}
	curScale := maxScale
	result[0] = curScale
	for depth := 1; depth < maxDepth; depth++ {
		curScale = curScale.Times(scale).F16()
	}
	return result
}

// Scales is a set of transformations which scale x and y centered at the
// origin.
type Scales [maxDepth]fixed.F16
