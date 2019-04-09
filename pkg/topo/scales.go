package topo

import (
	"willbeason/worldproc/pkg/fixed"
)

func PowerScales(maxScale, scale fixed.F16) Scales {
	result := Scales{}
	curScale := maxScale
	result[0] = curScale
	for depth := 1; depth < maxDepth; depth++ {
		curScale = curScale.Times(scale).F16()
	}
	return result
}

type Scales [maxDepth]fixed.F16
