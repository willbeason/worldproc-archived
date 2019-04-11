package transforms

import (
	"willbeason/worldproc/pkg/fixed"
)

const (
	// maxDepth of 50 supports scales of down to 1 pixel for relative scales of 0.8.
	// Relative scales more frequent than that are too noisy.
	maxDepth = 50
)

// Scales is a set of transformations which scale x and y centered at the
// origin.
type Scales [maxDepth]scale

type scale struct {
	Frequency, Amplitude fixed.F16
}

// PowerScales generates a set of scales where the factor from each scale to its
// neighbors is scale, with maxScale as the first.
func PowerScales(maxAmplitude, scale float64) Scales {
	result := Scales{}

	curAmplitude := maxAmplitude
	result[0].Amplitude = fixed.Float(curAmplitude)
	curFrequency := 1.0 / maxAmplitude
	result[0].Frequency = fixed.Float(curFrequency)
	for depth := 1; depth < maxDepth; depth++ {
		curAmplitude = curAmplitude * scale
		result[depth].Amplitude = fixed.Float(curAmplitude)
		curFrequency = curFrequency / scale
		result[depth].Frequency = fixed.Float(curFrequency)
	}
	return result
}
