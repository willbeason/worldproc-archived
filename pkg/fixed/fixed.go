// Package fixed is a performant library for non-integral math with 16 or 32 bits of precision.
package fixed

const (
	// size16 is the number of bits past the decimal.
	size16 = 16

	// Zero32 is the value 0 for use as a compile-time constant.
	Zero32 = F32(0)

	// One16 is the value 1 for use as a compile-time constant.
	One16 = F16(1 << size16)

	floatFactor = float64(1 << size16)
	invFloatFactor = float64(1.0 / floatFactor)

	// remainderMask16 provides a convenient value to bitwise-and with to get the non-integral
	// part of an F16.
	remainderMask16 = One16 - 1
)

// F16 represents nonnegative integral multiples of 2^-16 from 0 to 2^48 - 2^-16.
type F16 uint64

// F32 represents nonnegative integral multiples of 2^-32 from 0 to 2^32 - 2^-32.
type F32 uint64

// F48 represents nonnegative integral multiples of 2^-48 from 0 to 2^16 - 2^-48.
type F48 uint64

// Int converts an Int into an F16.
func Int(i int) F16 {
	return F16(i << size16)
}

// Float converts a float into an F16.
func Float(f float64) F16 {
	return F16(int(f * floatFactor))
}

// Times multiplies two F16s together exactly, returning an F32.
// This eliminates unnecessary bit-shifting on intermediate values.
func (f F16) Times(f2 F16) F32 {
	return F32(f * f2)
}

// Int returns the integral part of the F16.
//
// Measured slower to manually inline.
func (f F16) Int() int {
	return int(f >> size16)
}

// Float returns an equivalent floating-point representation of the F16.
func (f F16) Float() float64 {
	return float64(f) * invFloatFactor
}

// Remainder returns the non-integral part of the F16.
//
// Measured faster to store result if replacing 3 or more uses.
func (f F16) Remainder() F16 {
	return f & remainderMask16
}

// F16 returns a truncated version of the F32.
func (f F32) F16() F16 {
	return F16(f >> size16)
}