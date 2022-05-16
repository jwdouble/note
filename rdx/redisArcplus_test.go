package rdx

import (
	"arctron.lib/rdx"
	"fmt"
	"testing"
	"time"
)

func Test_Fuzzing(t *testing.T) {
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

	cmd := cli.Keys(ctx, "jw*")
	rows, _ := cmd.Result()
	for k, v := range rows {
		fmt.Println(k, "-->", v)
	}
}
