package GBK2UTF8

import (
	"fmt"
	"github.com/zhangyiming748/GBK2UTF8/log"
	"github.com/zhangyiming748/GBK2UTF8/mahonia"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"unicode/utf8"
)

func GBK2UTF8(src, pattern, dst string) {
	files := getFiles(src, pattern)
	for index, file := range files {
		runtime.GC()
		gbk2utf8_help(src, file, dst, index+1, len(files))
	}
}

func gbk2utf8_help(src, file, dst string, index, total int) {
	in := strings.Join([]string{src, file}, "/")
	extname := path.Ext(file)
	filename := strings.Trim(file, extname)
	filename = replace(filename)
	out := strings.Join([]string{dst, strings.Join([]string{filename, extname}, "")}, "/")
	if isUTF8(in) {
		writeUTF8(out, readUTF8(in))
		log.Debug.Printf("跳过编码已经是UTF8的文件:%s\n", src)
	} else {
		u8 := readGB18030(in)
		nums := writeUTF8(out, u8)
		log.Info.Printf("文件:%s写入%d个字符\n", out, nums)
	}
	log.Debug.Printf("处理完成第%d/%d个文件:%s\n", index, total, file)
}

func getFiles(dir, pattern string) []string {
	files, _ := os.ReadDir(dir)
	var aim []string
	types := strings.Split(pattern, ";") //"wmv;rm"
	for _, f := range files {
		if l := strings.Split(f.Name(), ".")[0]; len(l) != 0 {
			for _, v := range types {
				if strings.HasSuffix(f.Name(), v) {
					log.Debug.Printf("有效的目标文件:%v\n", f.Name())

					aim = append(aim, f.Name())
				}
			}
		}
	}
	return aim
}

func replace(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "《", "", -1)
	str = strings.Replace(str, "》", "", -1)
	str = strings.Replace(str, "【", "", -1)
	str = strings.Replace(str, "】", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, "+", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\u00A0", "", -1)
	str = strings.Replace(str, "\u0000", "", -1)
	return str
}

func isUTF8(src string) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	file, err := ioutil.ReadFile(src)
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
		log.Info.Println("编码不存在")
	}
	return decoder.ConvertString(string(file))
}

func readUTF8(src string) string {
	file, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func writeUTF8(dst, s string) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	f, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Debug.Println(err)
	}

	writeString, err := f.WriteString(s)
	if err != nil {
		panic(err)
	}
	return writeString
}
