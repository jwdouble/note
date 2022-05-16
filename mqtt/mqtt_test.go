package mqtt

import (
	"fmt"
	"testing"

	"arctron.lib/mqttx/v2"
)

func init() {
	mqttx.Register("tcp://root:123@219.146.74.162:1883")
}

func Sub() {
	err := mqttx.GetInstance().Subscribe("/application/alarm/#", 1, func(client mqttx.Client, message mqttx.Message) {
		fmt.Println("access#")
		fmt.Println(string(message.Payload()))
	})
	if err != nil {
		fmt.Println(err)
	}
	select {}
}

func Test_sub(t *testing.T) {
	fmt.Println("start")
	Sub()
}
