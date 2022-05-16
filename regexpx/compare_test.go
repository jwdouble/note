package regexpx

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func Test_compare(t *testing.T) {
	buf := "xxx<nihao>xxx<haha>xxx"
	reg := regexp.MustCompile(`<[A-Za-z]*>`)
	if reg != nil {
		s := reg.FindAllString(buf, -1)
		fmt.Println(s) //[3.14 345.12 7.8]
		ss := reg.FindAllStringSubmatch(buf, -1)
		fmt.Println(ss) //[[3.14] [345.12] [7.8]]
	}
}

func Test_compare1(t *testing.T) {
	var s []string
	buf := "qwe<nihao>asd<haha>zxc"
	reg := regexp.MustCompile(`>[A-Za-z]+`)
	if reg != nil {
		s = reg.FindAllString(buf, -1)
		fmt.Println(s) //[3.14 345.12 7.8]
		ss := reg.FindAllStringSubmatch(buf, -1)
		fmt.Println(ss) //[[3.14] [345.12] [7.8]]
	}

	var prefix []string
	n := strings.Index(buf, "<")
	prefix = append(prefix, buf[:n])
	for _, v := range s {
		prefix = append(prefix, strings.TrimPrefix(v, ">"))
	}
	fmt.Println(prefix)
}
