package syncx

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
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

func Test_Gnum(t *testing.T) {
	go func() {
		for {
			fmt.Println("G: ", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(100 * time.Millisecond)
	}

}
