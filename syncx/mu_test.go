package syncx

import (
	"fmt"
	"sync"
	"testing"

	"arctron.lib/codec/json"
)

func TestMu(t *testing.T) {
	// 锁：在同一时间只有一个进程可以占有锁，在对共享变量进行处理时使用锁就不会有竞争问题 -- 前提时所有对该数据的操作都要加锁

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

func TestMM(t *testing.T) {
	t1 := &MM{}
	t2 := &MM{}
	t1.handle()
	t2.handle() //  panic 两个相同的锁不能操作同一个全局变量
}

func Test_whetherNeedLock(t *testing.T) {
	type A struct {
		a string
	}
	type B struct {
		b string
	}
	type C struct {
		*A
		*B
	}

	a := &A{}
	b := &B{}
	c := &C{
		A: a,
		B: b,
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		a.a = "1"
		wg.Done()
	}()

	go func() {
		b.b = "2"
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(json.MustMarshalString(a.a), b.b, c) //yes
}
