package benchmark

import (
	"bytes"
	"strings"
	"testing"
	"unsafe"
)

//fmt.Sprintf 使用反射极消耗资源不适合高性能场景
//字符串的简单拼接用“+”就可以，如果是超长连续拼接则极耗资源
//bytes.write 速度比 strings.write 更快。在使用不安全指针的情况下 bytes + unsafe pointer 性能最高
//如果能丢到一个数组里慢慢join， join效率最高,不管是速度还是内存

func StringAdd() {
	var a string
	for i := 0; i < 1000; i++ {
		a += "adasdsafasfasdgasdgfa;iosfyhoqwhnflaskdhflakbfasudgflqkwjebflaiusgdfblkajsbzxlugbalfblqiauwegf"
	}
}

func StringBuilder() {
	var str strings.Builder
	for i := 0; i < 1000; i++ {
		str.WriteString("adasdsafasfasdgasdgfa;iosfyhoqwhnflaskdhflakbfasudgflqkwjebflaiusgdfblkajsbzxlugbalfblqiauwegf")
	}
	a := str.String()
	_ = a
}

func StringJoin() {
	var a string
	var elems []string
	for i := 0; i < 1000; i++ {
		elems = append(elems, "adasdsafasfasdgasdgfa;iosfyhoqwhnflaskdhflakbfasudgflqkwjebflaiusgdfblkajsbzxlugbalfblqiauwegf")
	}
	strings.Join(elems, a)
}

func BufWrite() {
	var buf bytes.Buffer
	for i := 0; i < 1000; i++ {
		buf.WriteString("adasdsafasfasdgasdgfa;iosfyhoqwhnflaskdhflakbfasudgflqkwjebflaiusgdfblkajsbzxlugbalfblqiauwegf")
	}
	a := buf.String()
	_ = a
}

func BufWrite1() {
	var buf bytes.Buffer
	for i := 0; i < 1000; i++ {
		buf.WriteString("adasdsafasfasdgasdgfa;iosfyhoqwhnflaskdhflakbfasudgflqkwjebflaiusgdfblkajsbzxlugbalfblqiauwegf")
	}
	a := *(*string)(unsafe.Pointer(&buf))
	_ = a
}

func BenchmarkStringAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringAdd()
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuilder()
	}
}

func BenchmarkBufWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BufWrite()
	}
}

func BenchmarkBufWrite1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BufWrite1()
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringJoin()
	}
}
