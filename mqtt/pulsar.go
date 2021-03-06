package mqtt

import (
	"context"
	"fmt"

	"arctron.lib/codec/json"
	"arctron.lib/conf"
	"arctron.lib/mq"
	"github.com/apache/pulsar-client-go/pulsar"
)

var client pulsar.Client

func init() {
	mq.Register(conf.APP_PULSAR_DSN.Value("pulsar://10.243.11.35:6650"))
}

func pub() {
	p, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "jwtest-1",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("send msg --------------------------------------------->")
	_, err = p.Send(context.Background(), &pulsar.ProducerMessage{Payload: []byte("hello")})
	if err != nil {
		fmt.Println(err)
	}
}

func read() {
	r, err := client.CreateReader(pulsar.ReaderOptions{
		Topic: "jwtest",
		Name:  "jw",
	})
	if err != nil {
		fmt.Println(err)
	}
	m, err := r.Next(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("----------->", string(m.Payload()))
}

func Listen1(ch chan string) {
	mq.GetInstance().ConsumeFunc("123-*", "123", func(ctx context.Context, msg mq.Messager) error {
		r := msg.Raw().(*mq.PulsarMessage).Topic()
		fmt.Println("listen1: ", r)
		ch <- r
		//fmt.Println(string(msg.Body()))
		return nil
	}, 5)
}

func Listen2() {
	mq.GetInstance().ConsumeFunc("456-*", "465", func(ctx context.Context, msg mq.Messager) error {
		fmt.Println("listen2: ", msg.Raw().(*mq.PulsarMessage).Topic())
		fmt.Println("")
		//fmt.Println(string(msg.Body()))
		return nil
	}, 5)
}

func Send1() {
	fmt.Println("send1")
	err := mq.GetInstance().Publish(context.Background(), "123-xxx11", json.MustMarshal("haha1"))
	if err != nil {
		fmt.Println(err)
	}
}

func Send2(ch chan string) {
	for v := range ch {
		fmt.Println("send2: ", v)
		err := mq.GetInstance().Publish(context.Background(), "456-xxx", json.MustMarshal("haha2"))
		if err != nil {
			fmt.Println(err)
		}
	}

}
