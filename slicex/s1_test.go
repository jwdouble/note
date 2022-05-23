package slicex

import (
	"fmt"
	"testing"
)

func Test_cal(t *testing.T) {
	fmt.Println(fmt.Sprintf("%s/business/workorder/getWorkOrderList.do?pageNo=1&pageSize=%d&dataSources=%s&projectId=%s&startTime=%s&endTime=%s", "www.baidu.com", 10, "haha", "1234", "0415", "0501"))
}

func Test_clean(t *testing.T) {
	s := []int{1, 2, 3}
	s = s[0:0]
	fmt.Println(s, len(s))
}
