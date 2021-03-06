package timex

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"arctron.lib/sqlkit"
	"arctron.lib/timex"
)

func Test_calculate(t *testing.T) {
	fmt.Println(time.Now().AddDate(0, 0, -24).Format(time.RFC3339))

	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().UnixMicro())
	fmt.Println(time.Now().UnixMilli())
	fmt.Println(strconv.Itoa(int(timex.BeginningOfDay().Unix())))
	t1 := time.Now()
	t2 := t1.AddDate(1, 1, 1)
	fmt.Println(t1, t2)
}

func Test_calculate1(t *testing.T) {
	c := time.Now()
	fmt.Println(time.Now().Unix())
	fmt.Println(sqlkit.TimestampNewLocal(c).GetSeconds())
}

func Test_sss(t *testing.T) {
	for {
		fmt.Println("1")
		time.Sleep(time.Duration(1 * time.Second))
	}
}
