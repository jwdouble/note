package mqtt

import (
	"arctron.lib/conf"
	"arctron.lib/mq"
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test_pulsar(t *testing.T) {
	fmt.Println("start")
	read()
	pub()
	fmt.Println("finish")
}

func Test_send(t *testing.T) {
	fmt.Println("send")
	Send()
}

func Test_listen(t *testing.T) {
	fmt.Println("listen")
	Listen()
	select {}
}

func Test_pn(t *testing.T) {
	go func() {
		for {
			fmt.Println("G: ", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		mq.GetInstance().ConsumeFunc("p1", "1", func(ctx context.Context, msg mq.Messager) error {
			return nil
		}, 100)
	}()

	mq.Register(conf.APP_PULSAR_DSN.Value("pulsar://10.243.11.35:6650"))
	ctx := context.Background()
	for {
		go func() {
			mq.GetInstance().Publish(ctx, "p1", []byte{1})
		}()

		time.Sleep(time.Second)
	}

}
