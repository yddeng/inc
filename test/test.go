package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Stdin.Close(); nil != err {
		fmt.Println(err)
		return
	}

	fmt.Println("sss")

	var s string
	fmt.Scan(&s)
	fmt.Println(s)
}
