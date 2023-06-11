package GBK2UTF8

import "testing"

func TestGBK2UTF8(t *testing.T) {
	src := "/Volumes/T7/slacking/Telegram/小说"
	//dst := "/Users/zen/Downloads/Telegram/Telegram/utf8"
	pattern := "txt;TXT"
	AllGBKs2UTF8(src, pattern)
}
