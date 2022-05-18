package timex

import (
	"arctron.lib/logx"
	"arctron.lib/timex"
	"fmt"
	"strings"
	"testing"
	"time"
)

func Test_timeParse(t *testing.T) {
	s := "2022-05-18 08:23:33"

	fmt.Println(time.ParseInLocation(timex.DateTimeFormat, s, time.Local))

}

func Test_timeFormat(t *testing.T) {
	current := time.Now()
	st := strings.Split(current.Format(timex.DateTimeFormat), " ")
	fmt.Println(st)
}

func Test_Puninx(t *testing.T) {
	fmt.Println(time.Unix(1651042518, 0))
	fmt.Println(time.Unix(time.Now().Unix(), 0))
	fmt.Println(time.Now().String())
}

func Test_info(t *testing.T) {
	logx.Warnf("haha")
	logx.Infof("haha")
}
