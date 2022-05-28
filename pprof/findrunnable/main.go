package main

import (
	"arctron.lib/conf"
	"arctron.lib/mq"
	"fmt"
	"golang.org/x/net/context"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	f, _ := os.Create("/tmp/cpu.pprof")
	pprof.StartCPUProfile(f)
	go func() {
		time.Sleep(10 * time.Second)
		pprof.StopCPUProfile()
		fmt.Println("record")
	}()

	mq.Register(conf.APP_PULSAR_DSN.Value("pulsar://10.243.11.35:6650"))

	go func() {
		//for {
		mq.GetInstance().Publish(context.Background(), "pprof-test", []byte("1"))
		time.Sleep(200 * time.Millisecond)
		//	}
	}()

	go func() {
		for {
			fmt.Println(runtime.NumGoroutine())
			time.Sleep(time.Second)
		}

		//fmt.Println(runtime.NumCPU())
	}()

	//T3()
	select {}
}

func T1() {
	fmt.Println("start")
	go T1Listen()
	go T1Send()

	select {}
}

var ch = make(chan int, 320)

func T1Listen() {
	var count int
	mq.GetInstance().ConsumeFunc("pprof-test", "note", func(ctx context.Context, msg mq.Messager) error {
		count++
		//fmt.Println(count)
		time.Sleep(1 * time.Microsecond)
		ch <- 1
		return nil
	}, 5)
}

func T1Send() {
	for c := range ch {
		_ = c
		go func() {
			for i := 0; i < 100000; i++ {
				_ = i
			}
		}()
	}
}

func T2() {
	go T2Listen()
	select {}
}

func T2Listen() {
	var count int
	mq.GetInstance().ConsumeFunc("pprof-test", "note", func(ctx context.Context, msg mq.Messager) error {
		count++
		//fmt.Println(count)
		time.Sleep(1 * time.Microsecond)
		for i := 0; i < 100000; i++ {
			_ = i
		}
		return nil
	}, 5)
}

func T3() {
	go T3Work()
	select {}
}

func T3Work() {
	select {
	case <-time.NewTicker(200 * time.Millisecond).C:
		func() {
			for i := 0; i < 100000; i++ {
				_ = i
			}
		}()
	}
}
