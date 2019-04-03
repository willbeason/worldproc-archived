package noise

import "willbeason/worldproc/pkg/fixed"

// Noise represents some form of noise to be used in generation.
//noinspection GoUnnecessarilyExportedIdentifiers
type Noise interface {
	// At returns the value of the noise at the passed position.
	At(x, y fixed.F16) fixed.F32
}
