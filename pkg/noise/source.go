package noise

import "willbeason/worldproc/pkg/fixed"

// Source represents a source of noise to be used in generation.
//noinspection GoUnnecessarilyExportedIdentifiers
type Source interface {
	// V returns the value of the noise at the passed position.
	V(x, y fixed.F16) fixed.F32
}
