package jsonx

import (
	"arctron.lib/codec/json"
	"arctron.lib/sqlkit"
	"fmt"
	"testing"
	"time"

	"arctron.lib/codec/proto"
)

func Test_proto1(t *testing.T) {
	m := map[string]interface{}{
		"1": "1",
	}

	fmt.Println(proto.MustNewProtoValue(m))

}

func Test_timestamp(t *testing.T) {
	type test struct {
		Ts *sqlkit.Timestamp `json:"ts"`
	}

	s := sqlkit.TimestampNewLocal(time.Now())

	a := fmt.Sprintf(`{"ts":"%s"}`, s.String())

	b := test{}

	err := json.Unmarshal([]byte(a), &b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}
