package mapx

import (
	"fmt"
	"testing"

	"arctron.lib/codec/json"
)

func Test_map(t *testing.T) {
	m := map[string]bool{
		"1": false,
		"2": true,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for n := range m { // print key
		fmt.Println(n)
	}
}

func Test_mapInit(t *testing.T) {
	m := map[int]int32{}
	for i := 0; i < 10; i++ {
		m[i]++ // yes
	}
	fmt.Println(m)
	fmt.Println(json.MustMarshalString(m))
}

func Test_mapLen(t *testing.T) {
	m := map[int]int32{}
	fmt.Println(len(m))
	m[1] = 1
	fmt.Println(len(m))
}
