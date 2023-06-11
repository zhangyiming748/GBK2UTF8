package GBK2UTF8

import (
	"fmt"
	"github.com/zhangyiming748/GBK2UTF8/mahonia"
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"golang.org/x/exp/slog"
	"os"
	"path"
	"strings"
	"unicode/utf8"
)

func AllGBKs2UTF8(root, pattern string) {
	folders := GetAllFolder.List(root)
	for _, folder := range folders {
		GBKs2UTF8(folder, pattern)
	}
}
func GBKs2UTF8(src, pattern string) {
	files := GetFileInfo.GetAllFileInfo(src, pattern)
	for index, file := range files {
		slog.Info(fmt.Sprintf("正在处理%d/%d个文件", index+1, len(files)))
		//runtime.GC()
		GBK2UTF8(file)
	}
}

func GBK2UTF8(info GetFileInfo.Info) {
	fp := info.FullPath
	prefix := path.Base(fp)
	newFp := strings.Join([]string{prefix, "utf8", info.FullName}, string(os.PathSeparator))
	fmt.Printf("in = %v\nout = %v\n", fp, newFp)
	if isUTF8(fp) {
		writeUTF8(newFp, readUTF8(fp))
		slog.Info("skip", slog.String("编码已经是UTF8,直接复制", info.FullName))
	} else {
		u8 := readGB18030(fp)
		nums := writeUTF8(newFp, u8)
		slog.Info("文件写入", slog.String("文件名", newFp), slog.Int("字符数", nums))
	}
}

func isUTF8(src string) bool {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("查询是否为UTF8时产生错误", slog.Any("错误文本", err))
		}
	}()
	file, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	return utf8.Valid(file)
}

func readGB18030(src string) string {
	file, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	decoder := mahonia.NewDecoder("gb18030")
	if decoder == nil {
		slog.Error("编码不存在", slog.Any("错误文本", err))
	}
	return decoder.ConvertString(string(file))
}

func readUTF8(src string) string {
	file, err := os.ReadFile(src)
	if err != nil {
		slog.Error("读取utf8产生错误", slog.Any("错误文本", err))
	}
	return string(file)
}

func writeUTF8(dst, s string) int {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("转写utf8产生错误", slog.Any("错误文本", err))
		}
	}()
	f, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		slog.Error("打开目标文件产生错误", slog.Any("错误文本", err))
	}

	writeString, err := f.WriteString(s)
	if err != nil {
		slog.Error("写文件产生错误", slog.Any("错误文本", err))
	}
	return writeString
}
