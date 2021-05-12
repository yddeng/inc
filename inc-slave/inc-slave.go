package main

import (
	"flag"
	"fmt"
	"github.com/yddeng/inc"
	"os"
)

func main() {
	fmt.Println(inc.Logo("inc-slave"))

	commandLine := flag.NewFlagSet("inc", flag.ExitOnError)
	a := commandLine.String("a", "", "--address     start server host, required ")
	n := commandLine.String("n", "slave", "--name     name for slave, optional ")
	commandLine.Parse(os.Args[1:])

	if *a == "" {
		return
	}

	_ = inc.LaunchIncSlave(*n, *a)

	select {}

}
