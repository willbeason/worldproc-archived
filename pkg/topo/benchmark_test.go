package topo_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"willbeason/worldproc/pkg/fixed"
	"willbeason/worldproc/pkg/topo"
)

var (
	src = rand.NewSource(time.Now().UnixNano())

	t = topo.Topography{
		Scales: topo.PowerScales(1000, 1.0 / math.Phi),
		Offsets: topo.RandomOffsets(src),
		Rotations: topo.RandomRotations(src),
		Depth: 24,
	}

	p1 = fixed.Float(rand.Float64() * 100)
	p2 = fixed.Float(rand.Float64() * 100)

	p3 fixed.F32
)

func init() {
	t.Noise.Fill(src)
	p3 = fixed.F32(p3)
}

func BenchmarkTopography_Height(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p3 = t.Height(p1, p2)
	}
}
