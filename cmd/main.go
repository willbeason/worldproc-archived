package main

import (
	"fmt"
	"math/rand"
	"time"

	"willbeason/worldproc/pkg/fixed"
	"willbeason/worldproc/pkg/noise"
)

func noise1() {
	var n = noise.Value{}
	n.Fill(rand.NewSource(time.Now().UnixNano()))

	start := time.Now()
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			n.At(fixed.Int(x), fixed.Int(y))
		}
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func noise10() {
	var n = noise.Value{}
	n.Fill(rand.NewSource(time.Now().UnixNano()))

	start := time.Now()
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			n.At(fixed.Int(x), fixed.Int(y))
			n.At(fixed.Int(x), fixed.Int(y))
			n.At(fixed.Int(x), fixed.Int(y))
			n.At(fixed.Int(x), fixed.Int(y))
			n.At(fixed.Int(x), fixed.Int(y))
			n.At(fixed.Int(x), fixed.Int(y))
			n.At(fixed.Int(x), fixed.Int(y))
			n.At(fixed.Int(x), fixed.Int(y))
			n.At(fixed.Int(x), fixed.Int(y))
			n.At(fixed.Int(x), fixed.Int(y))
		}
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func main() {
	noise1()
	noise10()
}
