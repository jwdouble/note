package structx

import (
	"time"
)

type demo struct {
	Name   string     `jsonx:"name"`
	Age    int        `jsonx:"age"`
	Record []struct { // 免定义结构体直接内部创建的方式
		Time    time.Time `jsonx:"timex"`
		Context string    `jsonx:"context"`
	} `jsonx:"record"`
}
