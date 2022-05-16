package proto

import (
	"arctron.os.api/common"
	"fmt"
	"testing"
)

func Test_protoVal(t *testing.T) {
	m := map[string]interface{}{
		"test": "11564fda4faf",
	}
	st := common.MustNewStruct(m)
	sv := common.NewStructValue(st)
	fmt.Println(sv.String())
	fmt.Println(sv.GetStructValue().Fields["test"].GetStringValue())
}
