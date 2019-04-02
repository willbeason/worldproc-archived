package noise

import "willbeason/worldproc/pkg/fixed"

const (
	// shift is a compile-time constant representing the number of bits to shift 1 by to get
	// the size of noise to use. We want the noise size to be a power of 2 so that we can
	// use bit shifts instead of multiplication and division, and bitwise operations for
	// moduli.
	shift = 7

	// size is the scale of noise in units (for certain implementations) before it repeats.
	// We have it as a compile-time constant so types can use it as array lengths.
	size = 1 << shift

	// noiseSize2 is the canonical square of size for use in 2D noise.
	noiseSize2 = size * size

	// intMask provides a convenient integer to take a bitwise-and with in order to perform
	// a cheap modulus.
	intMask = size - 1
)

// Noise represents some form of noise to be used in generation.
//noinspection GoUnnecessarilyExportedIdentifiers
type Noise interface {
	// At returns the value of the noise at the passed position.
	At(x, y fixed.F16) fixed.F32
}
