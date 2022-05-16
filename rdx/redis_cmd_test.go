package rdx

import (
	"arctron.lib/rdx"
	"fmt"
	"testing"
	"time"
)

func Test_Scan(t *testing.T) {
	cli := rdx.GetInstance()

	cli.GetWithFetch(ctx, "jw:asd1", func() (string, time.Duration, error) {
		return "haha1", time.Minute, nil
	})
	cli.GetWithFetch(ctx, "jw:asd2", func() (string, time.Duration, error) {
		return "haha1", time.Minute, nil
	})
	cli.GetWithFetch(ctx, "jw:asd3:haha", func() (string, time.Duration, error) {
		return "haha1", time.Minute, nil
	})

	//
	cmd := cli.Scan(ctx, 0, "jw*", 10) //scan命令： cursor 游标：从第几个数据开始 match匹配规则 count扫描多少个数据
	rows, n, _ := cmd.Result()
	fmt.Println(n)
	for k, v := range rows {
		fmt.Println(k, "-->", v)
	}
}
