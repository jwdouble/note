package mqtt

import (
	"fmt"
	"strconv"
	"time"

	"arctron.lib/conf"
	"arctron.lib/mqttx/v2"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
var c mqtt.Client

//func init() {
//	mqtt.DEBUG = log.New(os.Stdout, "", 0)
//	mqtt.ERROR = log.New(os.Stdout, "", 0)
//	opts := mqtt.NewClientOptions().
//		AddBroker("10.243.11.35:1883").
//		SetUsername("arc-os").
//		SetPassword("p1ssw0rd").
//		SetClientID("11")
//
//	opts.SetKeepAlive(5 * timex.Second)
//	// Set the message callback handler
//	opts.SetDefaultPublishHandler(f)
//	opts.SetPingTimeout(100 * timex.Second)
//
//	c = mqtt.NewClient(opts)
//	if token := c.Connect(); token.Wait() && token.Error() != nil {
//		panic(token.Error())
//	}
//
//}

func emqMain() {
	// Subscribe to a topic
	c.Subscribe("m", 1, nil)

	// Publish a message
	go func() {
		for i := 0; i < 10; i++ {
			// $delayed/{DelayInterval}/{TopicName}
			token := c.Publish("$delayed/10/m", 1, false, "----> jw"+strconv.Itoa(i))
			token.Wait()
			time.Sleep(time.Millisecond * 10)
		}
	}()

	//timex.Sleep(10 * timex.Second)
	//
	//// Unscribe
	//if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
	//	fmt.Println(token.Error())
	//	os.Exit(1)
	//}

	time.Sleep(10 * time.Second)

	// Disconnect
	//c.Disconnect(250)
	time.Sleep(1 * time.Second)
}

func emqByGoLib() {
	mqttx.Register(conf.APP_MQTT_DSN.Value("tcp://arc-os:p1ssw0rd@10.243.11.35:1883"))
	_ = mqttx.GetInstance().
		Publish("$delayed/3/meetingSch/", 1, false, "fuck")
}
