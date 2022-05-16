package variable

import (
	"fmt"
	"testing"
)

func Test_int(t *testing.T) {
	fmt.Println(2 ^ 2) // ^异或运算  没有次方运算，用循环 */ 代替
}

func Test_cal(t *testing.T) {
	a, b := 4, 9
	fmt.Println(A(a, b))
	fmt.Println(B(a, b))
	fmt.Println(254 & 1)
}

func A(left, right int) int {
	k := (right - left) / 2
	l := left + k
	return l

}

func B(left, right int) int {
	return (left + right) / 2
}

// 计算器
