package GBK2UTF8

import "testing"

func TestGBK2UTF8(t *testing.T) {
	src := "/Users/zen/Documents/未编码"
	//dst := "/Users/zen/Downloads/Telegram/Telegram/utf8"
	pattern := "txt;TXT"
	GBKs2UTF8(src, pattern)

}
