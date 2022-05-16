package rdx

import (
	"fmt"
	"testing"
	"time"

	"jw.lib/conf"
	"jw.lib/rdx"
)

func Test_redis1(t *testing.T) {
	rdx.Register(conf.AppRedisConn.Value(rdx.DefaultRedisAddr))
	cli := rdx.GetRdxOperator()
	cli.Set("test", "hahha", time.Hour)
	//fmt.Println(cli.Get("test"))
}

func Test_redis1ListAll(t *testing.T) {
	rdx.Register(conf.AppRedisConn.Value(rdx.DefaultRedisAddr))
	cli := rdx.GetRdxOperator()
	s := cli.Keys("*")
	fmt.Println(s)

	for {
		fmt.Println("--->", cli.LPop("logx").String())
		fmt.Println("lpop logx: redis: nil" == cli.LPop("logx").String())
		time.Sleep(time.Millisecond * 500)
	}

}
