package runtimex

import (
	"fmt"
	"runtime"
	"testing"
)

func Test_caller(t *testing.T) {
	A()
}

func A() {
	B()
}

func B() {
	C()
}

func C() {
	fmt.Println(runtime.Caller(0))
	fmt.Println(runtime.Caller(1))
	fmt.Println(runtime.Caller(2))
	fmt.Println(runtime.Caller(3))
}
