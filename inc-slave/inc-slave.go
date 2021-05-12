package main

import (
	"flag"
	"fmt"
	"github.com/yddeng/inc"
	"os"
)

func logo() string {
	l := `
___  ________   _________  ___  ___  ________      
|\  \|\   ___  \|\___   ___\\  \|\  \|\   ___  \        Internal Network Tunnel  
\ \  \ \  \\ \  \|___ \  \_\ \  \\\  \ \  \\ \  \       
 \ \  \ \  \\ \  \   \ \  \ \ \  \\\  \ \  \\ \  \      LEAF
  \ \  \ \  \\ \  \   \ \  \ \ \  \\\  \ \  \\ \  \     
   \ \__\ \__\\ \__\   \ \__\ \ \_______\ \__\\ \__\
    \|__|\|__| \|__|    \|__|  \|_______|\|__| \|__|      
`

	return l
}

func main() {
	fmt.Println(logo())

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
