package chx

import (
	"fmt"
	"testing"
	"time"
)

func Test_chInterface(t *testing.T) {
	ch := make(chan interface{}, 1)
	ch <- 1
	fmt.Println(<-ch)
	ch <- "haha"
	fmt.Println(<-ch)
	fmt.Println(ch)
}

func Test_chRange(t *testing.T) {
	ch := make(chan int, 10)
	go func() {
		for {
			select {
			case <-time.Tick(1 * time.Second):
				ch <- 1
			}
		}
	}()

	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1

	for v := range ch {
		fmt.Println(v)
	}

}
