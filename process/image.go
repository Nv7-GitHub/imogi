package main

import (
	"image"

	"github.com/lucasb-eyer/go-colorful"
)

const min = 0.9 // Minimum percentage of solid colors

func isOk(image image.Image) (bool, [3]int) {
	dim := image.Bounds()

	clearCount := 0
	for x := 0; x < dim.Dx(); x++ {
		for y := 0; y < dim.Dy(); y++ {
			pix := image.At(x, y)
			_, _, _, a := pix.RGBA()
			if a < 50 {
				clearCount++
			}
		}
	}

	percClear := float64(clearCount) / float64(dim.Dx()*dim.Dy())
	if percClear > 1-min {
		return false, [3]int{}
	}

	// Get average color
	r := 0
	g := 0
	b := 0
	cnt := 0
	for x := 0; x < dim.Dx(); x++ {
		for y := 0; y < dim.Dy(); y++ {
			pix := image.At(x, y)
			rV, gV, bV, a := pix.RGBA()
			rV, gV, bV = rV>>8, gV>>8, bV>>8 // convert to out of 255
			if a > 50 {
				r += int(rV)
				g += int(gV)
				b += int(bV)
				cnt++
			}
		}
	}
	r = r / cnt
	b = b / cnt
	g = g / cnt
	color := colorful.Color{R: float64(r) / float64(255), G: float64(g) / float64(255), B: float64(b) / float64(255)}
	h, s, v := color.Hsv()

	return true, [3]int{int(h), int(s * 255), int(v * 255)}
}
