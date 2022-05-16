package reflectx

import (
	"arctron.lib/field"
	"fmt"
	"testing"
)

type A struct {
	s string
	k string
}

func Test_lib(t *testing.T) {
	var aa []*A
	aa = append(aa, &A{s: ""})
	ans := field.Get(aa, "s")
	if ans != nil { // panic的原因是对nil断言，先判断nil就行了
		a := ans.([]string)
		fmt.Println(a)
	}
}
