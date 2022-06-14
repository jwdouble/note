package error

import (
	"fmt"
	"testing"

	"arctron.lib/errx/v2"
)

func Test_errPrint(t *testing.T) {
	e := errx.New(404, 100, "test")
	fmt.Println(fmt.Sprintf("kankan: %s", e))
}

func Test_fmt(t *testing.T) {
	s := []interface{}{"1", "2", 3}
	fmt.Println(fmt.Sprintf("a: %s b: %s c: %d", s...))
}
