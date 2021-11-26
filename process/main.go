package main

import (
	"encoding/json"
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/schollz/progressbar/v3"
)

type Emoji struct {
	Text  string
	Color [3]int // LAB
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	emojis, err := os.ReadDir("emojis")
	handle(err)

	out := make([]Emoji, 0)
	bar := progressbar.New(len(emojis))

	for _, file := range emojis {
		// Read img
		f, err := os.Open("emojis/" + file.Name())
		handle(err)
		img, err := png.Decode(f)
		handle(err)
		f.Close()

		ok, color := isOk(img)
		if ok {
			ext := filepath.Ext(file.Name())
			name := file.Name()[:len(file.Name())-len(ext)]
			emoji := Emoji{
				Text:  filenameToString(name),
				Color: color,
			}
			out = append(out, emoji)
		}

		bar.Add(1)
	}
	bar.Finish()
	fmt.Println()

	// Save
	outF, err := os.Create("../emojis.json")
	handle(err)
	defer outF.Close()

	enc := json.NewEncoder(outF)
	err = enc.Encode(out)
	handle(err)
}

func filenameToString(filename string) string {
	codepoints := strings.Split(filename, "-")
	chars := make([]rune, len(codepoints))
	for i, codepoint := range codepoints {
		val, err := strconv.ParseInt(codepoint, 16, 32)
		handle(err)

		chars[i] = rune(val)
	}
	return string(chars)
}
