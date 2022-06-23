package syncx

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 锁：在同一时间只有一个进程可以占有锁，在对共享变量进行处理时使用锁就不会有竞争问题 ***前提是所有对该数据的操作都要加锁***
//
func Test_Mu1(t *testing.T) { // 没有操作都加锁导致的加锁无效
	var i int
	mu := sync.Mutex{}
	go func() {
		for {
			mu.Lock()
			fmt.Println(i)
			time.Sleep(time.Second * 10)
			mu.Unlock()
		}
	}()

	for i < 100 {
		i++
		time.Sleep(time.Second)
	}
}

func Test_Mu2(t *testing.T) { // 加锁成功但是开销巨大
	var i int
	mu := sync.Mutex{}
	go func() {
		for {
			mu.Lock()
			fmt.Println(i)
			time.Sleep(time.Second * 10)
			mu.Unlock()
		}
	}()

	for i < 100 {
		mu.Lock()
		i++
		mu.Unlock()
		time.Sleep(time.Second)
	}
}

func Test_Mu3(t *testing.T) { // 使用原子操作
	var i int32

	go func() {
		for {
			val := atomic.LoadInt32(&i)
			fmt.Println(val)
			time.Sleep(time.Second * 5)
		}
	}()

	for i < 100 {
		atomic.AddInt32(&i, 1)
		time.Sleep(time.Second)
	}
}

var num = 1

type MM struct {
	mu sync.Mutex
}

func (m *MM) handle() {
	m.mu.Lock()
	num++
	m.mu.Unlock()
}

// panic 两个相同的锁不能操作同一个全局变量
func TestMM(t *testing.T) {
	t1 := &MM{}
	t2 := &MM{}
	t1.handle()
	t2.handle()
}
