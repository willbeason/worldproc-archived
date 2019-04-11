package fixed_test

import (
	"math/rand"
	"testing"
	"time"

	"willbeason/worldproc/pkg/fixed"
)

// exponentBits is the number of bits a float64 uses to represent its exponent.
const exponentBits = 10

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestF16_FloatEquivalent(t *testing.T) {
	// Tests when F16.Float64() guarantees it returns an equivalent float64.
	for i := 0; i < 100; i++ {
		x := r.Uint64()
		x >>= exponentBits + 1
		//   64 bits generated
		// - 10 bits lost to give space for the exponent
		// -  1 bit  lost to rounding
		// 53 significant bits are exact.

		f16 := fixed.F16(x)
		l := f16.Float64()
		if fixed.Float(l) != f16 {
			t.Fatalf("%v is not equal to %v", f16, fixed.Float(l))
		}
	}
}

func TestF16_FloatTruncate(t *testing.T) {
	// Tests when F16.Float64() guarantees it returns a nearly equivalent float64.

	// Casting to a float64 loses up to 10 digits of precision.
	// At most this round trip has a relative error of 2^-53. (63-10=53)
	// Worst case is 2^63+2^10-1 gets rounded to 2^63.
	maxDiff := 1 << exponentBits

	for i := 0; i < 100; i++ {
		x := r.Uint64()

		f16 := fixed.F16(x)
		l := f16.Float64()

		diff := int(uint(f16) - uint(fixed.Float(l)))
		if diff < 0 {
			diff = -diff
		}

		if diff > maxDiff {
			t.Fatalf("|%v - %v| = %v > %v", f16, fixed.Float(l), diff, maxDiff)
		}
	}
}
