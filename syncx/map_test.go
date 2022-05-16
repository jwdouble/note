package syncx

import (
	"fmt"
	"sync"
	"testing"
)

func Test_syncMap(t *testing.T) {
	m := sync.Map{}

	m.LoadOrStore("k", "v")

	fmt.Println(m.Load("k"))
}
