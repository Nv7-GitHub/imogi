package main

import (
	"image"

	"github.com/lucasb-eyer/go-colorful"
)

const min = 0.8        // Minimum percentage of solid colors
const maxHSVDiff = 800 // Maximum difference in HSV
const maxLabDiff = 1   // Maximum difference in LAB

func isOk(image image.Image) (bool, [3]int) {
	dim := image.Bounds()

	clearCount := 0
	for x := 0; x < dim.Dx(); x++ {
		for y := 0; y < dim.Dy(); y++ {
			pix := image.At(x, y)
			_, _, _, a := pix.RGBA()
			a = a >> 8
			if a < 50 {
				clearCount++
			}
		}
	}

	percClear := float64(clearCount) / float64(dim.Dx()*dim.Dy())
	if percClear > 1-min {
		return false, [3]int{}
	}

	// Get average color & diff
	l := float64(0)
	a := float64(0)
	b := float64(0)
	cnt := float64(0)
	lowestHSV := float64(821)
	highestHSV := float64(-821)
	lowestLAB := float64(201)
	highestLAB := float64(-201)
	for x := 0; x < dim.Dx(); x++ {
		for y := 0; y < dim.Dy(); y++ {
			pix := image.At(x, y)
			rV, gV, bV, aV := pix.RGBA()
			rV, gV, bV, aV = rV>>8, gV>>8, bV>>8, aV>>8 // convert to out of 255
			if aV > 50 {
				col := colorful.Color{R: float64(rV) / 255, G: float64(gV) / 255, B: float64(bV) / 255}
				lV, aV, bV := col.Lab()
				l += lV
				a += aV
				b += bV

				// Diff HSV
				h, s, v := col.Hsv()
				sum := h + s*360 + v*360
				if sum < lowestHSV {
					lowestHSV = sum
				}
				if sum > highestHSV {
					highestHSV = sum
				}
				cnt++

				// Diff LAB
				sum = lV + aV + bV
				if sum < lowestLAB {
					lowestLAB = sum
				}
				if sum > highestLAB {
					highestLAB = sum
				}
			}
		}
	}

	// check contrast
	if highestHSV-lowestHSV > maxHSVDiff || highestLAB-lowestLAB > maxLabDiff {
		return false, [3]int{}
	}

	l = l / cnt
	a = a / cnt
	b = b / cnt

	return true, [3]int{int(l * 255), int(a * 255), int(b * 255)}
}
