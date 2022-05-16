package error

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

type ErrTest struct {
	str string
}

func (e *ErrTest) Error() string {
	return e.str
}

// error本质是个接口，对于底层抛上来的error可以尝试断言后取值使用
//例：func IsErrDuplicate(err error) bool {
//     if e, ok := err.(*pgconn.PgError); ok {
//        return e.Code == "23505"
//     }
//     return false
//    }

// error只和他本身相等

func Test(t *testing.T) {
	fmt.Println(errors.New("1") == errors.New("1")) //false
}
