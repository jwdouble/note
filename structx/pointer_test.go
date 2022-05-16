package structx

import (
	"fmt"
	"testing"
)

type Info struct {
	a string
}

func Test_pointer(t *testing.T) {
	fmt.Println(r1())
	fmt.Println(r2())
}

func r1() []Info {
	var list []Info
	list = append(list, Info{a: "1"})
	list = append(list, Info{a: "2"})

	return list
}

func r2() []*Info {
	var list []*Info
	list = append(list, &Info{a: "1"})
	list = append(list, &Info{a: "2"})

	return list
}
