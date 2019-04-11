package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"

	"willbeason/worldproc/pkg/fixed"
	"willbeason/worldproc/pkg/topo"
)

var (
	phone = image.Rect(0, 0, 750, 1334)

	sz = phone
)

func main() {
	img := image.NewRGBA(sz)

	src := rand.NewSource(time.Now().UnixNano())

	t := topo.Topography{
		Scales: topo.PowerScales(200, 1.0 / math.SqrtPhi),
		Offsets: topo.RandomOffsets(src),
		Rotations: topo.RandomRotations(src),
		Depth: 20,
	}
	t.Noise.Fill(src)

	factor := 256.0 / (4676.0 / 5)

	for x := 0; x < sz.Max.X; x++ {
		for y := 0; y < sz.Max.Y; y++ {
			h := t.Height(fixed.Int(x), fixed.Int(y))
			i := uint8(h.Float() * factor)
			img.Set(x, y, color.RGBA{
				R: i,
				G: i,
				B: i,
				A: 255,
			})
		}
	}

	fw, _ := os.OpenFile("C:\\Users\\Will\\Pictures\\WorldProc\\img.png", os.O_CREATE | os.O_WRONLY, os.ModePerm)

	_ = png.Encode(fw, img)


}
