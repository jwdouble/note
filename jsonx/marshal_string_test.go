package jsonx

import (
	"arctron.lib/codec/json"
	"fmt"
	"testing"
)

func Test_string(t *testing.T) {
	var resp []string
	rs(&resp)

	fmt.Println(resp)
}

func rs(x interface{}) {
	s := []string{"1", "2"}
	buf := json.MustMarshal(s)
	json.Unmarshal(buf, x)
}
