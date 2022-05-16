package mqtt

import (
	"fmt"
	"testing"
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
