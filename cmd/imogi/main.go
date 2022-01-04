package main

import (
	"fmt"
	"image"
	"os"
	"strings"

	_ "embed"
	_ "image/jpeg"
	_ "image/png"

	"github.com/Nv7-Github/imogi"
	"github.com/nfnt/resize"
)

const charsPerMsg = 2000

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	inp := GetInput("Name of image: ")
	f, err := os.Open(inp)
	handle(err)
	defer f.Close()

	img, _, err := image.Decode(f)
	handle(err)

	width := GetInputInt("Max Width (0 for unlimited, recommend 14-30): ")
	height := GetInputInt("Max Height (0 for unlimited, usually leave at 0): ")
	img = resize.Resize(uint(width), uint(height), img, resize.NearestNeighbor)

	/*// Save resized
	outF, err := os.Create("out.png")
	handle(err)
	defer outF.Close()
	err = png.Encode(outF, img)
	handle(err)*/

	// Convert
	out := imogi.ConvertImg(img)

	// Get msgs
	msgs := make([]string, 0)
	cnt := 0
	txt := ""
	for _, line := range strings.Split(out, "\n") {
		// Calculate size
		length := 0
		for _, r := range line {
			shortcut, exists := imogi.GetEmojiShortcut(r)
			if exists {
				length += len(shortcut) + 2 // (2 for the two ":"s)
			} else {
				length += 10 // arbitrary number
				fmt.Println("Doesnt exist", string(r))
			}
		}

		// Split
		if cnt+length > charsPerMsg {
			msgs = append(msgs, txt)
			txt = ""
			cnt = 0
		}
		cnt += length
		txt += line
		txt += "\n"
	}

	// Print with line breaks for discord
	fmt.Printf("There are %d chunks. Copy the following chunks into messages in discord, one by one. Press enter to go to the next chunk.\n", len(msgs))
	for _, msg := range msgs {
		reader.ReadLine()

		fmt.Println(msg)
		fmt.Println()
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Finished!")
}
