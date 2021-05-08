package main

import (
	"fmt"
	"os"
	"strings"
)

func readLine() string {
	buffer := make([]byte, 128)
	n, _ := os.Stdin.Read(buffer)
	return string(buffer[:n-1])
}

func main() {
	fmt.Println("exe1")

	s := " df fff  fdsg   sg    fd "

	s = strings.(s, "  ", " ")
	fmt.Println(s)

	for {
		line := readLine()
		switch line {
		case "quit":
			fmt.Println("quit")
			return
		default:
			fmt.Println("-->", line)
		}
	}
}
