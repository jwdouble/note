package benchmark

import "testing"

// make 和 append 法相比，make初始化速度更快，内存分配更小，内存分配次数更少

func StructSlice1() {
	var a []*Middle
	for i := 0; i < 10; i++ {
		a = append(a, &Middle{})
	}
}

func StructSlice2() {
	a := make([]*Middle, 10)
	for i := 0; i < 10; i++ {
		a[i] = &Middle{}
	}
}

func BenchmarkStructSlice1(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StructSlice1()
	}
}

func BenchmarkStructSlice2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StructSlice2()
	}
}
