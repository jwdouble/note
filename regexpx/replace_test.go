package regexpx

import (
	"fmt"
	"regexp"
	"testing"
)

//func replace (s string) string{
//
//}

func Test_replace(t *testing.T) {
	data := "aaa!<type>bbb!<name>ccc!<url>ddd"

	//通过函数进行替换
	re3, _ := regexp.Compile(`!<[A-Za-z]*>`)

	//把匹配的所有字符a替换成b
	rep2 := re3.ReplaceAllString(data, "b")
	fmt.Println(rep2)
}
