package jsonx

import (
	"fmt"
	"testing"

	"arctron.lib/codec/json"
)

type Hum struct {
	A string `json:"a"`
	S string `json:"s"`
}

func Test_new(t *testing.T) {
	h := Hum{
		A: "aa",
		S: "ss",
	}

	buf := json.MustMarshal(&h)
	fmt.Println(string(buf))
}
