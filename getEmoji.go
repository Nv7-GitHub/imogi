package imogi

import (
	"embed"
	"image"
	"image/png"
	"strconv"
)

//go:embed emojis/*.png
var emojiVal embed.FS

func GetEmoji(emoji rune) (image.Image, error) {
	hex := strconv.FormatInt(int64(emoji), 16)
	f, err := emojiVal.Open(hex + ".png")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return png.Decode(f)
}
