package jsonx

import (
	"fmt"
	"testing"

	"jw.lib/jsonx"
)

func Test_xx1(t *testing.T) {
	i := 135469873134736873
	r := jsonx.MustMarshal(i)
	var s string
	jsonx.MustUnmarshal(r, &s)
	fmt.Println(s)
}
