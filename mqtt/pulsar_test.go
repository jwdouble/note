package mqtt

import (
	"fmt"
	"testing"
	"time"
)

func Test_pulsar(t *testing.T) {
	fmt.Println("start")
	read()
	pub()
	fmt.Println("finish")
}

//func Test_send(t *testing.T) {
//	Send1()
//}

func Test_listen(t *testing.T) {
	fmt.Println("listen")
	ch1 := make(chan string, 32)
	go Send2(ch1)
	Send1()
	time.Sleep(time.Millisecond * 10)
	Listen1(ch1)
	Listen2()

	select {}
}

func Test_sendx(t *testing.T) {
	Send1()
}
