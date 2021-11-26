package main

import (
	"encoding/json"
	"fmt"
	"image/png"
	"os"
	"strings"

	_ "embed"

	"github.com/nfnt/resize"
)

//go:embed emojis.json
var emojiData string

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

type Emoji struct {
	Text  string
	Color [3]int // LAB
}

var emojis []Emoji

func main() {
	err := json.Unmarshal([]byte(emojiData), &emojis)
	handle(err)

	f, err := os.Open("image.png")
	handle(err)
	defer f.Close()

	img, err := png.Decode(f)
	handle(err)

	// Resize to 10x10
	img = resize.Resize(25, 25, img, resize.NearestNeighbor)

	// Save resized
	outF, err := os.Create("out.png")
	handle(err)
	defer outF.Close()
	err = png.Encode(outF, img)
	handle(err)

	// Convert
	out := convertImg(img)

	// Print with line breaks for discord
	cnt := 0
	for _, line := range strings.Split(out, "\n") {
		length := len([]byte(line))
		if cnt+length > 1000 {
			fmt.Println()
			fmt.Println()
			cnt = 0
		}
		cnt += length
		fmt.Println(line)
	}
}
