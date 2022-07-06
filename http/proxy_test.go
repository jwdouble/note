package http

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"testing"
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
	b := make([]byte, 1024)

	n, err := conn.Read(b)
	fmt.Println(n, err)

	buf := bytes.NewBuffer(b)
	data, err := buf.ReadBytes('\n')
	pushErr(err)

	fmt.Println(data)

	list := strings.Fields(string(data))
	method := list[0]
	url := list[1]
	fmt.Println(method, url, list)

	resp, err := net.Dial("tcp", "150.158.7.96:80")
	pushErr(err)
	resp.Read(b)
	conn.Write(b)
}
