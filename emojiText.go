package imogi

import (
	_ "embed"
	"encoding/json"
)

//go:embed emojiNames.json
var emojiNames []byte

var emojiText map[rune]string

func init() {
	dat := make(map[string]string)
	json.Unmarshal(emojiNames, &dat)

	emojiText = make(map[rune]string, len(dat))
	for k, v := range dat {
		emojiText[[]rune(k)[0]] = v
	}
}

func GetEmojiShortcut(emoji rune) (string, bool) {
	v, exists := emojiText[emoji]
	return v, exists
}
