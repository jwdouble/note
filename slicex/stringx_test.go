package slicex

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func Test_stringReplace(t *testing.T) {

}

func insertData(src string, m map[string]string) (tar string) {
	for k, v := range m {
		src = strings.Replace(src, "${"+k+"}", v, -1)
	}

	tar = src

	return
}

func Test_stringRange(t *testing.T) {
	s := "1234567"
	fmt.Println(s[0])
	// s[0] = s[2] 	字符串不可修改
}

func Test_stringChinese(t *testing.T) {
	length := 0
	s := "中国123456789"
	var buf []rune
	var flagHan bool
	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			flagHan = true
			fmt.Println(length, "-->", string(v))
		} else {
			fmt.Println(length, "-->", string(v))
		}
		if flagHan && length == 10 {
			break
		}
		buf = append(buf, v)
		length++
	}
	fmt.Println(string(buf))
}

func trimAllSpace(s string) string {
	var buf []rune
	for _, v := range s {
		if v != 32 {
			buf = append(buf, v)
		}
	}
	return string(buf)
}

func Test_trimString(t *testing.T) {
	s := "  a v c d   a  "
	fmt.Println(s)
	s = strings.TrimSpace(s)
	fmt.Println(s)
	s = trimAllSpace(s)
	fmt.Println(s)
}
