package syncx

import (
	"fmt"
	"sync"
	"testing"
)

func Test_gofunc(t *testing.T) {
	a := 1
	wg := sync.WaitGroup{}

	wg.Add(1) // yes
	go func() {
		// wg.Add(1)  -- 放在里面协程会秒退 no
		defer wg.Done()
		a++
		fmt.Println("++")
	}()

	wg.Wait()
	fmt.Println(a)
}
