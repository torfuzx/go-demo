package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "[::1]:12800")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dialer := &net.Dialer{
		LocalAddr: addr,
	}

	conn, err := dialer.Dial("tcp", ":9009")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("conn is: %#v\n", conn.LocalAddr())

	fmt.Printf("dialer is: %#v\n", dialer)
}
