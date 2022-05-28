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
		Topic: "jwtest",
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

func Listen() {
	mq.GetInstance().ConsumeFunc("event-*-*", "cmpt-event", func(ctx context.Context, msg mq.Messager) error {
		fmt.Println(msg.Body())
		return nil
	}, 1)
}

func Send() {
	err := mq.GetInstance().Publish(context.Background(), "event-alarmNotice-quvrkqwy6fs7posx", json.MustMarshal("haha"))
	if err != nil {
		fmt.Println(err)
	}
}
