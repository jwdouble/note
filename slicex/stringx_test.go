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

func Test_split(t *testing.T) {
	s := `${aa}${bb}${cc}`
	fmt.Println(strings.Split(s, `${`))
	s1 := "abcdefg"
	fmt.Println(s1)
	r := strings.Split(s1, "")
	fmt.Println(r, len(r))
}

func renderDataSM(src string, m map[string]string) string {
	var param []string
	sp := strings.Split(src, `${`)
	for n := range sp {
		i := strings.IndexRune(sp[n], '}')
		if i > 0 {
			param = append(param, sp[n][:i])
		}
	}

	var data strings.Builder
	data.WriteString(`[`)
	for n, v := range param {
		if n != 0 {
			data.WriteString(`,`)
		}
		data.WriteString(`"`)
		data.WriteString(m[v])
		data.WriteString(`"`)
	}
	data.WriteString(`]`)

	return data.String()
}

func Test_rand(t *testing.T) {
	r := renderDataSM(`haha${aa}xx${bb}888`, map[string]string{
		"aa": "1",
		"bb": "2",
	})
	fmt.Println(r)
}

func Test_getBigTypeCode(t *testing.T) {
	r := getBigTypeCode("30-50.66.77")
	fmt.Println(r)
	r1 := getBigTypeCode("30-50.66.77.88.99")
	fmt.Println(r1)
	r2 := getBigTypeCode("30-50.66")
	fmt.Println(r2)
}

func getBigTypeCode(code string) string {
	l := strings.Split(code, ".")
	if len(l) >= 3 {
		l = l[:1]
		l = append(l, []string{"00", "00"}...)
	} else {
		return ""
	}
	return strings.Join(l, ".")
}
