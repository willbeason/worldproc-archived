package topo

import (
	"willbeason/worldproc/pkg/fixed"
)

// PowerScales generates a set of scales where the factor from each scale to its
// neighbors is scale, with maxScale as the first.
func PowerScales(maxAmplitude, scale float64) Scales {
	result := Scales{}

	curAmplitude := maxAmplitude
	result.amplitude[0] = fixed.Float(curAmplitude)
	curFrequency := 1.0 / maxAmplitude
	result.frequency[0] = fixed.Float(curFrequency)
	for depth := 1; depth < maxDepth; depth++ {
		curAmplitude = curAmplitude * scale
		result.amplitude[depth] = fixed.Float(curAmplitude)
		curFrequency = curFrequency / scale
		result.frequency[depth] = fixed.Float(curFrequency)
	}
	return result
}

// Scales is a set of transformations which scale x and y centered at the
// origin.
type Scales struct {
	frequency [maxDepth]fixed.F16
	amplitude [maxDepth]fixed.F16
}
