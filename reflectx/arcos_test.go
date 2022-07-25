package reflectx

import (
	"fmt"
	"reflect"
	"testing"

	"arctron.lib/field"
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

func Test_reflectx(t *testing.T) {
	aa := &A{
		s: "1",
		k: "2",
	}
	v := reflect.ValueOf(aa)
	fmt.Println(v.Kind())
	fmt.Println(v.Type())
	fmt.Println(v.Elem())
	fmt.Println(v.Field(0))
}
