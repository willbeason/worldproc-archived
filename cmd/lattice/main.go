package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"willbeason/worldproc/pkg/colors"
	"willbeason/worldproc/pkg/fixed"
	"willbeason/worldproc/pkg/lattice"
	"willbeason/worldproc/pkg/noise"
	"willbeason/worldproc/pkg/topo"
	"willbeason/worldproc/pkg/transforms"
)

const (
	ytCellSize = 130
	invYtCellSize = 1.0 / ytCellSize
)

var (
	blog    = image.Rect(0, 0, 300, 300)
	youTube = image.Rect(0, 0, 1280, 720)

	sz = youTube

	dir = fmt.Sprintf("C:\\Users\\Will\\Pictures\\WorldProc\\%d", time.Now().Unix())

	src = rand.NewSource(time.Now().UnixNano())

	topography = topo.NoiseTopography{
		Noise: noise.Value{},
		Scales: transforms.PowerScales(ytCellSize, 1.0/math.SqrtPhi),
		Offsets: transforms.RandomOffsets(src),
		Rotations: transforms.RandomRotations(src),
		Depth: 1,
	}
)

func forPix(rect image.Rectangle, f func(x, y int) color.Color) *image.NRGBA {
	img := image.NRGBA{
		Pix:     make([]uint8, rect.Max.X*rect.Max.Y*4),
		Stride:  rect.Max.X*4,
		Rect:    rect,
	}

	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			c := f(x, y)
			img.Set(x, y, c)
		}
	}

	return &img
}

func blend(img1, img2 *image.NRGBA, p func(x, y int) float64) *image.NRGBA {
	return forPix(img1.Rect, func(x, y int) color.Color {
		return colors.Blend(img1.NRGBAAt(x, y), img2.NRGBAAt(x, y), p(x, y))
	})
}

func heightToColorIndex(tallest float64) func(height float64) float64 {
	if tallest == 0.0 {
		return func(height float64) float64 {
			return 0.0
		}
	}

	tallestInv := 1.0 / tallest
	deep := 0.50
	shallow := 0.55
	sand := 0.56
	grass := 0.60
	forest := 0.70
	deepForest := 0.80
	stone := 0.90
	snow := 0.95

	return func(height float64) float64 {
		f := height * tallestInv
		switch {
		case f < deep:
			return 0.0
		case f <= shallow:
			return (f - deep)*(20.0)
		case f <= sand:
			return (f - shallow)*(100.0) + 1.0
		case f <= grass:
			return (f - sand)*(25.0) + 2.0
		case f <= forest:
			return (f - grass)*(10.0) + 3.0
		case f <= deepForest:
			return (f - forest)*(10.0) + 4.0
		case f <= stone:
			return (f - deepForest)*(10.0) + 5.0
		case f <= snow:
			return (f - stone)*(20.0) + 6.0
		}
		return 7.0
	}
}

func adj(rect image.Rectangle, x, y int) (float64, float64) {
	return float64(x-rect.Max.X/2)*invYtCellSize, float64(y-rect.Max.Y/2)*invYtCellSize
}

func main() {
	topography.Noise.Fill(src)
	topography.Offsets[0] = transforms.Offset(fixed.Float(float64(-sz.Max.X/2)*invYtCellSize), fixed.Float(float64(-sz.Max.Y/2)*invYtCellSize))
	topography.Rotations[0] = transforms.NoRotation

	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	topography.Noise.Fill(rand.NewSource(time.Now().UnixNano()))

	nt := topo.NewTopography(sz, func(x, y int) fixed.F32 {
		return topography.HeightNearest(
			fixed.Int(x),
			fixed.Int(y),
		)
	})
	lt := topo.NewTopography(sz, func(x, y int) fixed.F32 {
		return topography.HeightLinear(
			fixed.Int(x),
			fixed.Int(y),
		)
	})

	l := lattice.NewLattice(fixed.Float(ytCellSize))

	white := forPix(sz, func(x, y int) color.Color {
		return color.White
	})

	delay := 20

	centerX := sz.Max.X / 2
	centerY := sz.Max.Y / 2
	for frame := 0; frame < 101; frame++ {
		latticeR := float64((frame - delay) * 20)

		curTop := topo.NewTopography(sz, func(x, y int) fixed.F32 {
			xp, yp := adj(sz, x, y)
			xp, yp = math.Round(xp), math.Round(yp)
			r := math.Sqrt(float64(xp*xp + yp*yp))*ytCellSize
			diff := latticeR - r

			if frame < 70 {
				factor := (diff + 200) / 400

				if factor < 0.0 {
					factor = 0.0
				} else if factor > 1.0 {
					factor = 1.0
				}

				return fixed.Float32(nt.Height(x, y).Float64() * factor)
			}

			factor := float64(frame - 70) / 30.0
			return fixed.Float32(nt.Height(x, y).Float64() * (1.0 - factor)) + fixed.Float32(lt.Height(x, y).Float64() * factor)
		})

		h2c := heightToColorIndex(1.0)

		curCol := forPix(sz, func(x, y int) color.Color {
			return colors.TopoScale.At(h2c(curTop.Height(x, y).Float64()))
		})

		img := blend(curCol, white, func(x, y int) float64 {
			xp, yp := adj(sz, x, y)
			r := math.Sqrt(xp*xp + yp*yp)*ytCellSize
			diff := (r - latticeR) / 50

			d := l.Dist(x-centerX+ytCellSize/2, y-centerY+ytCellSize/2).Float64()
			return math.Exp(-2000 * d * d - diff * diff) / 2.0
		})

		fw, err := os.OpenFile(filepath.Join(dir, fmt.Sprintf("out-%03d.png", frame)), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = png.Encode(fw, img)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}


