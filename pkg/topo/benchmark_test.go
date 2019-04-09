package topo_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"willbeason/worldproc/pkg/fixed"
	"willbeason/worldproc/pkg/noise"
	"willbeason/worldproc/pkg/topo"
)

var (
	src = rand.NewSource(time.Now().UnixNano())

	scales = topo.PowerScales(fixed.Int(1000), fixed.Float(1.0 / math.Phi))
	offsets = topo.RandomOffsets(src)
	rotations = topo.RandomRotations(src)

	t = topo.ValueNoise{
		Noise: noise.Value{},
		Scales: scales,
		Offsets: offsets,
		Rotations: rotations,
		Depth: 25,
	}

	p1 = fixed.Float(rand.Float64() * 100)
	p2 = fixed.Float(rand.Float64() * 100)

	p4 fixed.F32
)

func init() {
	t.Noise.Fill(src)
	p4 = fixed.F32(p4)
}

func BenchmarkTopograhpy_Height(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p4 = t.Height(p1, p2)
	}
}
