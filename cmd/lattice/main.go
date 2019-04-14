package main

import (
	"fmt"
	"image"
	"image/gif"
	"math/rand"
	"os"
	"time"

	"willbeason/worldproc/pkg/color"
	"willbeason/worldproc/pkg/fixed"
	"willbeason/worldproc/pkg/lattice"
	"willbeason/worldproc/pkg/noise"
)

const (
	nFrames  = 102
	cellSize = 29.89
)

var (
	blog    = image.Rect(0, 0, 300, 300)
	youTube = image.Rect(0, 0, 1280, 720)

	sz = blog
	n  = noise.Value{}
)

func forPix(rect image.Rectangle, f func(x, y int) uint8) *image.Paletted {
	img := image.Paletted{
		Pix:     make([]uint8, rect.Max.X*rect.Max.Y),
		Stride:  rect.Max.X,
		Rect:    rect,
		Palette: color.Grayscale(),
	}

	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			img.Pix[x+img.Stride*y] = f(x, y)
		}
	}

	return &img
}

func main() {
	n.Fill(rand.NewSource(time.Now().UnixNano()))

	genLattice()
	genValues()
	genInterpolate()
	genYouTube()
}

func genLattice() {
	fw, _ := os.OpenFile("C:\\Users\\Will\\Pictures\\WorldProc\\value_noise_1.gif", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	g := gif.GIF{
		Image:    make([]*image.Paletted, nFrames),
		Delay:    make([]int, nFrames),
		Disposal: make([]byte, nFrames),
	}

	baseL := lattice.NewLattice(fixed.Float(cellSize * 10))
	g.Image[0] = forPix(sz, func(x, y int) uint8 {
		h := baseL.V(fixed.Int(x), fixed.Int(y)).Int()
		if h == 1 {
			return 1
		}
		return 255
	})
	g.Delay[0] = 50

	maxCol := 200

	for frame := 1; frame < 101; frame++ {
		l := lattice.NewLattice(fixed.Float(cellSize))

		g.Delay[frame] = 4
		g.Disposal[frame] = gif.DisposalPrevious

		tFrame := frame - 1

		g.Image[frame] = forPix(sz, func(x, y int) uint8 {
			if baseL.V(fixed.Int(x), fixed.Int(y)).Int() == 1 {
				return 0
			}

			val := x*x + y*y
			maxVal := tFrame * tFrame * 36
			if val > maxVal {
				return 0
			}

			prevVal := (tFrame - 1) * (tFrame - 1) * 36

			pct := float64(maxVal-val) / float64(maxVal-prevVal)
			if pct > 1.0 {
				pct = 1.0
			}

			h := l.V(fixed.Int(x), fixed.Int(y)).Int()
			if h == 1 {
				return 255 - uint8(float64(maxCol)*pct)
			}
			return 0
		})
	}

	g.Image[101] = forPix(sz, func(x, y int) uint8 {
		return 0
	})
	g.Delay[100] = 1000
	g.Disposal[101] = gif.DisposalPrevious

	err := gif.EncodeAll(fw, &g)
	if err != nil {
		fmt.Println(err)
	}

	_ = fw.Close()

}

func genValues() {
	fw, _ := os.OpenFile("C:\\Users\\Will\\Pictures\\WorldProc\\value_noise_2.gif", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	g := gif.GIF{
		Image:    make([]*image.Paletted, nFrames),
		Delay:    make([]int, nFrames),
		Disposal: make([]byte, nFrames),
	}

	baseL := lattice.NewLattice(fixed.Float(cellSize))
	g.Image[0] = forPix(sz, func(x, y int) uint8 {
		h := baseL.V(fixed.Int(x), fixed.Int(y)).Int()
		if h == 1 {
			return 1
		}
		return 255
	})
	g.Delay[0] = 50

	for frame := 1; frame < 101; frame++ {

		g.Delay[frame] = 4
		g.Disposal[frame] = gif.DisposalPrevious

		tFrame := frame - 1

		g.Image[frame] = forPix(sz, func(x, y int) uint8 {
			if baseL.V(fixed.Int(x), fixed.Int(y)).Int() == 1 {
				// Never overwrite grid.
				return 0
			}

			val := n.V(fixed.Int(int(float64(x)/cellSize)), fixed.Int(int(float64(y)/cellSize)))

			f := 1.0 - (1.0-val.Float64())*float64(tFrame)/100.0
			p := uint8(f * 255)
			if p == 0 {
				p = 1
			}
			return uint8(p)
		})
	}

	g.Image[101] = forPix(sz, func(x, y int) uint8 {
		return 0
	})
	g.Delay[100] = 1000
	g.Disposal[101] = gif.DisposalPrevious

	err := gif.EncodeAll(fw, &g)
	if err != nil {
		fmt.Println(err)
	}

	_ = fw.Close()

}

func genInterpolate() {
	fw, _ := os.OpenFile("C:\\Users\\Will\\Pictures\\WorldProc\\value_noise_3.gif", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	g := gif.GIF{
		Image:    make([]*image.Paletted, nFrames-50),
		Delay:    make([]int, nFrames-50),
		Disposal: make([]byte, nFrames-50),
	}

	baseL := lattice.NewLattice(fixed.Float(cellSize))
	g.Image[0] = forPix(sz, func(x, y int) uint8 {
		var f float64
		if baseL.V(fixed.Int(x), fixed.Int(y)).Int() == 0 {
			// Slowly overwrite grid.
			f = n.V(fixed.Int(int(float64(x)/cellSize)), fixed.Int(int(float64(y)/cellSize))).Float64()
		}

		p := uint8(f * 255)
		if p == 0 {
			p = 1
		}
		return p
	})
	g.Delay[0] = 100

	for frame := 1; frame < 51; frame++ {

		g.Delay[frame] = 4
		g.Disposal[frame] = gif.DisposalPrevious

		tFrame := frame - 1

		g.Image[frame] = forPix(sz, func(x, y int) uint8 {
			var approx float64
			if baseL.V(fixed.Int(x), fixed.Int(y)).Int() == 0 {
				// Slowly overwrite grid.
				approx = n.V(fixed.Int(int(float64(x)/cellSize)), fixed.Int(int(float64(y)/cellSize))).Float64()
			}
			val := n.V(fixed.Float(float64(x)/cellSize)-fixed.Float(0.5), fixed.Float(float64(y)/cellSize)-fixed.Float(0.5)).Float64()

			f := approx*(1.0-(float64(tFrame)/50.0)) + val*(float64(tFrame)/50.0)
			p := uint8(f * 255)
			if p == 0 {
				p = 1
			}
			return p
		})
	}

	g.Image[51] = forPix(sz, func(x, y int) uint8 {
		return 0
	})
	g.Delay[50] = 500
	g.Disposal[51] = gif.DisposalPrevious

	err := gif.EncodeAll(fw, &g)
	if err != nil {
		fmt.Println(err)
	}

	_ = fw.Close()

}

func genYouTube() {
	fw, _ := os.OpenFile("C:\\Users\\Will\\Pictures\\WorldProc\\value_noise_you_tube.gif", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	g := gif.GIF{
		Image:    make([]*image.Paletted, 500),
		Delay:    make([]int, 500),
		Disposal: make([]byte, 500),
	}

	ytCellSize := 39.9444
	baseL := lattice.NewLattice(fixed.Float(ytCellSize))
	g.Image[0] = forPix(youTube, func(x, y int) uint8 {
		return 255
	})
	g.Delay[0] = 5

	for frame := 1; frame < 101; frame++ {
		g.Delay[frame] = 5
		g.Disposal[frame] = gif.DisposalPrevious

		tFrame := frame - 1
		g.Image[frame] = forPix(youTube, func(x, y int) uint8 {
			val := x*x + y*y
			maxVal := tFrame * tFrame * 288
			if val > maxVal {
				return 0
			}
			prevVal := (tFrame - 1) * (tFrame - 1) * 288

			pct := float64(maxVal-val) / float64(maxVal-prevVal)
			if pct > 1.0 {
				pct = 1.0
			}
			h := baseL.V(fixed.Int(x), fixed.Int(y)).Int()
			if h == 1 {
				return 255 - uint8(float64(254)*pct)
			}
			return 0
		})
	}

	for frame := 101; frame < 201; frame++ {
		g.Delay[frame] = 5
		g.Disposal[frame] = gif.DisposalPrevious

		tFrame := frame - 101
		g.Image[frame] = forPix(youTube, func(x, y int) uint8 {
			if baseL.V(fixed.Int(x), fixed.Int(y)).Int() == 1 {
				// Never overwrite grid.
				return 1
			}

			val := n.V(fixed.Int(int(float64(x)/ytCellSize)), fixed.Int(int(float64(y)/ytCellSize)))

			f := 1.0 - (1.0-val.Float64())*float64(tFrame)/100.0
			p := uint8(f * 255)
			if p == 0 {
				p = 1
			}
			return uint8(p)
		})
	}

	for frame := 201; frame < 301; frame++ {
		g.Delay[frame] = 5
		g.Disposal[frame] = gif.DisposalPrevious

		tFrame := frame - 201
		g.Image[frame] = forPix(youTube, func(x, y int) uint8 {
			var approx float64
			if baseL.V(fixed.Int(x), fixed.Int(y)).Int() == 0 {
				// Slowly overwrite grid.
				approx = n.V(fixed.Int(int(float64(x)/ytCellSize)), fixed.Int(int(float64(y)/ytCellSize))).Float64()
			}
			val := n.V(fixed.Float(float64(x)/ytCellSize)-fixed.Float(0.5), fixed.Float(float64(y)/ytCellSize)-fixed.Float(0.5)).Float64()

			f := approx*(1.0-(float64(tFrame)/100.0)) + val*(float64(tFrame)/100.0)
			p := uint8(f * 255)
			if p == 0 {
				p = 1
			}
			return p
		})
	}

	for frame := 301; frame < 500; frame++ {
		g.Delay[frame] = 5
		g.Disposal[frame] = gif.DisposalPrevious
		g.Image[frame] = forPix(youTube, func(x, y int) uint8 {
			f := n.V(fixed.Float(float64(x)/ytCellSize)-fixed.Float(0.5), fixed.Float(float64(y)/ytCellSize)-fixed.Float(0.5)).Float64()
			p := uint8(f * 255)
			if p == 0 {
				p = 1
			}
			return p
		})
	}

	err := gif.EncodeAll(fw, &g)
	if err != nil {
		fmt.Println(err)
	}

	_ = fw.Close()

}
