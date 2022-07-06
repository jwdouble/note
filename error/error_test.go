package error

import (
	"fmt"
	"testing"

	"arctron.lib/errx/v2"
	"arctron.lib/logx"
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

func throwErr(err error, funcName string) {
	if err != nil {
		logx.WithCaller().Errorf(funcName+" %s", err)
	}
}

func Test_x1(t *testing.T) {

	err := errx.New(9999, 9999, "test err")

	throwErr(err, "nn")

	logx.WithCaller().Errorf("err: 2")

	logx.WithCaller().Errorf("err: 3")
}
