package syncx

import (
	"fmt"
	"sync"
	"testing"
)

func Test_pool(t *testing.T) {
	p := sync.Pool{
		New: func() interface{} {
			a := "init"
			return a
		},
	}

	p.Put("haha")
	fmt.Println(p.Get())
	p.Put("hehe")
	p.Put("hehe1")
	fmt.Println(p.Get())
}
