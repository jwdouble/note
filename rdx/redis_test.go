package rdx

import (
	"context"
	"fmt"
	"testing"
	"time"

	"arctron.lib/codec/json"
	"arctron.lib/conf"
	"arctron.lib/rdx"
)

var ctx = context.Background()

func init() {
	rdx.Register(conf.APP_REDIS_DSN.Value("redis://:p1ssw0rd@10.243.11.35:6379"))
}

func Test_GetWithFetch(t *testing.T) {
	r, err := rdx.GetInstance().GetWithFetch(ctx, "jwtest", func() (string, time.Duration, error) {
		return "hahah1", time.Hour, nil
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
	type data struct {
		Name string
		Val  string
	}
	req := data{
		Name: "jw1",
		Val:  "hhh",
	}

	buf, err := json.Marshal(&req) // 对结构体指针和结构体取地址的区别
	if err != nil {
		fmt.Println(err)
	}

	resp := rdx.GetInstance().Set(ctx, "jwtest", buf, time.Minute)
	fmt.Println(resp)
}

func Test_increase(t *testing.T) {
	num, err := rdx.GetInstance().GetWithFetch(ctx, "test_increase", func() (string, time.Duration, error) {
		return "1", time.Hour, nil
	})
	fmt.Println(num, err)

	rdx.GetInstance().Incr(ctx, "test_increase")

	num, err = rdx.GetInstance().GetWithFetch(ctx, "test_increase", nil)
	fmt.Println(num, err)
}

func Test_increase1(t *testing.T) {
	rdx.GetInstance().Incr(ctx, "test_increase1")

	num, err := rdx.GetInstance().GetWithFetch(ctx, "test_increase1", nil)
	fmt.Println(num, err)
}
