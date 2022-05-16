package timex

import (
	"fmt"
	"testing"
	"time"
)

const layout = "2006-01-02 15:04:05"

func Test_timeLocal(t *testing.T) {
	//默认时区
	testTime := "1995-08-04 02:05:05"

	fmt.Println(time.Now())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().UTC().Local())
	fmt.Println(time.Now().Local().Local())
	fmt.Println(time.Now().UTC().UTC())
	// result: timex now return timex with local timex zone

	fmt.Println("/**********************************/")
	fmt.Println(time.Parse(layout, testTime))
	fmt.Println(time.ParseInLocation(layout, testTime, time.UTC))
	fmt.Println(time.ParseInLocation(layout, testTime, time.Local))
	// result: timex parse return utc,recommend use parse in location

	fmt.Println("/**********************************/")
	current := time.Now()
	fmt.Println(time.Unix(current.Unix(), 0))
	fmt.Println(time.Unix(current.UTC().Unix(), 0))
	// result: timex unix return timex with local zone

	fmt.Println(time.Unix(0, 0))
	fmt.Println(time.Unix(0, 0).UTC())

	// 数据库输入输出对时区的影响
	// result: same with sqlx zone
}

func Test_parse(t *testing.T) {
	str := "2016/0/29"
	c, _ := time.ParseInLocation("2006/1/2", str, time.Local) // yes!
	fmt.Println(c)
}
