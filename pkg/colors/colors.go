package colors

import (
	"image/color"
)

var (
	DeepWater = color.NRGBA{R: 20, G: 20, B: 60, A: 255}
	ShallowWater = color.NRGBA{R: 57, G: 212, B: 243, A: 255}
	Sand = color.NRGBA{R: 194, G: 178, B: 128, A: 255}
	Grass = color.NRGBA{R: 114, G: 143, B: 81, A: 255}
	Forest = color.NRGBA{R: 34, G: 108, B: 34, A: 255}
	DeepForest = color.NRGBA{R: 14, G: 68, B: 14, A: 255}
	Stone = color.NRGBA{R: 139, G: 141, B: 122, A: 255}
	Snow = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

	TopoScale = Scale{
		colors: []color.NRGBA{
			DeepWater, // 0
 			ShallowWater, // 1
			Sand, // 2
			Grass, // 3
			Forest, // 4
			DeepForest, // 5
			Stone, // 6
			Snow, // 7
		},
	}
)
