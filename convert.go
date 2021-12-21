package imogi

import (
	"image"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

func ConvertImg(img image.Image) string {
	out := &strings.Builder{}
	bounds := img.Bounds()
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			r, g, b, a := img.At(x, y).RGBA()
			a = a >> 8
			if a < 50 {
				out.WriteString("⬛")
			} else {
				r, g, b = r>>8, g>>8, b>>8
				color := colorful.Color{R: float64(r) / 255, G: float64(g) / 255, B: float64(b) / 255}
				l, a, b := color.Lab()
				out.WriteString(getEmoji([3]int{int(l * 255), int(a * 255), int(b * 255)}))
			}
		}
		out.WriteString("\n")
	}
	return out.String()
}

// color is hsv
func getEmoji(color [3]int) string {
	dist := getDist(color, emojis[0].Color)
	text := emojis[0].Text
	for _, emoji := range emojis[1:] {
		distV := getDist(color, emoji.Color)
		if distV < dist {
			dist = distV
			text = emoji.Text
		}
	}
	return text
}

func getDist(val1 [3]int, val2 [3]int) int {
	return abs(val1[0]-val2[0]) + abs(val1[1]-val2[1]) + abs(val1[2]-val2[2])
}

func abs(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}
