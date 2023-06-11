package GBK2UTF8

import "testing"

func TestGBK2UTF8(t *testing.T) {
	src := "G:\\slacking\\Telegram\\小说\\未整理\\未编码"
	//dst := "/Users/zen/Downloads/Telegram/Telegram/utf8"
	pattern := "txt"
	GBKs2UTF8(src, pattern)

}
