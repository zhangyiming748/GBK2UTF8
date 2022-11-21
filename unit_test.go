package GBK2UTF8

import "testing"

func TestGBK2UTF8(t *testing.T) {
	src := "/Users/zen/Downloads/Telegram/Telegram"
	dst := "/Users/zen/Downloads/Telegram/Telegram/utf8"
	pattern := "txt"
	GBK2UTF8(src, pattern, dst)

}
