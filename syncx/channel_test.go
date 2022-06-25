package syncx

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//
func Test_c(t *testing.T) {
	go func() {
		for {
			fmt.Println(runtime.NumGoroutine())
			time.Sleep(200 * time.Millisecond)
		}
	}()

	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		time.Sleep(400 * time.Millisecond)
	}
}
