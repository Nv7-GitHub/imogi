package imogi

import (
	_ "embed"
	"encoding/json"
)

//go:embed emojis.json
var emojiData string

type Emoji struct {
	Text  string
	Color [3]int // LAB
}

var emojis []Emoji

func init() {
	err := json.Unmarshal([]byte(emojiData), &emojis)
	if err != nil {
		panic(err)
	}
}
