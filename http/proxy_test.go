package http

import (
	"fmt"
	"net"
	"strconv"
	"testing"

	"jw.lib/jsonx"
)

func pushErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Test_httpProxy(t *testing.T) {
	lis, err := net.Listen("tcp", ":8080")
	pushErr(err)

	for {
		conn, err := lis.Accept()
		pushErr(err)

		go handle(conn)
	}

}

func handle(conn net.Conn) {
	buf := make([]byte, 1024*10)

	n, err := conn.Read(buf[:])
	pushErr(err)

	type resp struct {
		Code int32  `json:"code,omitempty"`
		Msg  string `json:"msg,omitempty"`
	}

	n, err = conn.Write(jsonx.MustMarshal(&resp{Code: 200, Msg: strconv.Itoa(n)}))
	fmt.Println(n, err)
}
