package timex

import (
	"fmt"
	"testing"
	"time"
)

func Test_timeParse(t *testing.T) {
	s := "2020-12-11T12:01:39.256 +08:00"

	fmt.Println(time.ParseInLocation("2006-01-02T15:04:05 +08:00", s, time.Local))

}

func Test_timeFormat(t *testing.T) {
	current := time.Now()
	fmt.Println(current)
	fmt.Println(current.Format("15:04"))
}

func Test_Puninx(t *testing.T) {
	fmt.Println(time.Unix(1651042518, 0))
	fmt.Println(time.Unix(time.Now().Unix(), 0))
	fmt.Println(time.Now().String())
}
