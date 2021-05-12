package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":2345")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(l.Addr(), l.Addr().String())

}
