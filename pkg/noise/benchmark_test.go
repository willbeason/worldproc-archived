package noise

import (
	"math/rand"
	"testing"
)

func randP() pos {
	return p(rand.Float64()*float64(size), rand.Float64()*float64(size))
}

var px = randP()

func BenchmarkValue_At(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n.At(px.x, px.y)
	}
}
