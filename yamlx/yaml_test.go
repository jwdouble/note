package yamlx

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_yaml(t *testing.T) {
	fmt.Println("start ------>")

	data, err := os.ReadFile("./test.yaml")
	if err != nil {
		fmt.Println(err)
	}

	sm := sync.Map{}
	buf := strings.Split(string(data), "\n")

	for n := range buf {
		buf[n] = strings.TrimSuffix(buf[n], ":")
	}

	for i := len(buf) - 1; i >= 0; i-- {
		str := TrimAllSpace(buf[i])

		cols := strings.Split(str, ":")
		if len(cols) > 1 {
			pwd := YamlPwd(buf[:i])
			sm.Store(pwd+"."+cols[0], cols[1])

		}
	}

	fmt.Println(sm.Load("app.name"))
	fmt.Println(sm.Load("cache.list"))
	fmt.Println(sm.Load("mysql.host"))
}

func YamlPwd(elems []string) string {
	per := 999
	var pwd []string
	for i := len(elems) - 1; i >= 0; i-- {
		n := findFirst(elems[i])
		if n < per {
			var temp []string
			pwd = append(temp, elems[i])
			pwd = append(temp, pwd...)
		}
		if n == 0 {
			break
		}
	}
	return strings.Join(pwd, ".")
}
