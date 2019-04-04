package fixed

import (
	"math"
	"math/rand"
	"testing"
)

var (
	f1 = rand.Float64() * 100
	f2 = rand.Float64() * 10

	f3 float64

	p1 = Float(rand.Float64() * 10)
	p2 = Float(rand.Float64() * 10)

	p3 F16
	p4 F32

	i3 int
)

func init() {
	f3 = f3
	p3 = p3
	p4 = p4
	i3 = i3
}

func BenchmarkFloat_Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f3 = f1 * f2
	}
}

func BenchmarkF16_Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p4 = p1.Times(p2)
	}
}

func BenchmarkFloat_Remainder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f3 = f1 - float64(int(f1))
	}
}

func BenchmarkFloat_Remainder2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f3 = f1 - math.Floor(f1)
	}
}

func BenchmarkF16_Remainder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p3 = p1.Remainder()
	}
}

func BenchmarkFloat_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i3 = int(f1)
	}
}

func BenchmarkF16_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i3 = p1.Int()
	}
}
