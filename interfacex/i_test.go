package interfacex

import (
	"fmt"
	"testing"
)

func Test_i_array(t *testing.T) {
	s := []string{"1", "2", "hha"}
	var x []interface{}
	for _, v := range s {
		x = append(x, v)
	}

	for _, v := range x {
		fmt.Println(v.(string))
	}

}
