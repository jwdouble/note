package timex

import (
	"fmt"
	"testing"
	"time"

	"arctron.lib/sqlkit"
)

func TestNull(t *testing.T) {
	var tt time.Time
	fmt.Println(tt)
	fmt.Println(tt.Unix())
	fmt.Println("---------------------------")
	tt = time.Unix(0, 0)
	fmt.Println(tt)
	fmt.Println(tt.Unix())

	fmt.Println("<--------------------------->")

	var ts sqlkit.Timestamp
	fmt.Println(ts.AsTime())
	fmt.Println(ts.Seconds)

	// result: time类型未初始化状态下不能使用， timestamp可以  默认值为 1970-01-01 00:00:00 +0000 即时间戳0
}

//

func Test_uninx(t *testing.T) {
	i := 0
	for {
		if i > 10 {
			return
		}
		fmt.Println(time.Now().Unix())
		i++
		time.Sleep(time.Second)
	}

}
