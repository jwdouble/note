package slicex

import (
	"fmt"
	"strings"
	"testing"

	"main.go/yamlx"
)

func Test_length(t *testing.T) {
	s := []int{1, 1, 1, 1, 1, 1, 1}
	fmt.Println(len(s))
	if len(s) > 6 {
		s = s[3:]
	}

	fmt.Println(len(s))
}

func Test_append(t *testing.T) {
	var a1 []string
	var a2 []string

	a1 = append(a1, a2...)

	fmt.Println(a1)
}

func Test_init(t *testing.T) {
	slc := make([]string, 2)
	fmt.Println(len(slc))
	if len(slc) > 100 && slc[3] == "" { //info: && 符号会先判断表达式左边，如果成立再判断表达式右边
		fmt.Println("1")
	}

	var s1 []string
	s2 := make([]string, 0)
	s1 = append(s1, "1")
	s2 = append(s2, "1")
	fmt.Println(s1, s2) // 两种初始化方式都不会创造 [0]

}

func Test_range(t *testing.T) {
	a := make([]int, 3)
	b := make([]int, 3)
	c := []int{1, 2, 3}
	for n, v := range c {
		b[0] = v
		a[n] = b[0]
	}
	fmt.Println(a)
}

func Test_trim(t *testing.T) {
	s := "  1 2 3  4  "
	fmt.Println(strings.Split(s, " "))
	fmt.Println(strings.Split(s, "  "))
	fmt.Println(strings.Trim(s, "  ")) // 多少个空格都一样 同下
	fmt.Println(strings.TrimSpace(s))  // 只能消除收尾的空格
	fmt.Println(yamlx.TrimAllSpace(s)) //yes

	s1 := "a:b"
	s2 := "a:"
	fmt.Println(strings.Split(s1, ":"), len(strings.Split(s1, ":")), strings.Split(s1, ":")[1])
	fmt.Println(strings.Split(s2, ":"), len(strings.Split(s2, ":")), strings.Split(s2, ":")[1])

	fmt.Println(strings.Split("abc!@#poi", "!@#")) // sep不显示
}

func Test_replace(t *testing.T) {
	m := map[string]string{}
	s := `{"id":"218611273284190335", "parentId":"1046604776407057", "name":"综合管理中心（深圳）", "meta":{"createBy":"arc-mgr", "createAt":"2021-08-26T03:07:52Z"}, "code":"Z0000178913A8C83206EEF56394168C7", "property":1}`
	//s = strings.Replace(s,` `,``,-1)
	//s = strings.Replace(s,`"`,``,-1)
	//s = strings.Replace(s,"{","",-1)
	//s= strings.Replace(s,"}","",-1)

	r := strings.NewReplacer(` `, ``, `"`, ``, `{`, ``, `}`, ``)
	s = r.Replace(s)
	str := strings.Split(s, ",")
	for _, v := range str {
		params := strings.Split(v, ":")
		m[params[0]] = params[1]
	}

	for k, v := range m {
		fmt.Println(k + "--->" + v)
	}
}

func Test_len(t *testing.T) {
	var a, b map[string]string
	var list []string
	l := len(list)
	fmt.Println(l, &list, &a, &b)
}
