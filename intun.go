package main

import (
	"flag"
	"fmt"
	"github.com/yddeng/intun/client"
	"github.com/yddeng/intun/master"
	"os"
)

func logo() string {
	l := `
___  ________   _________  ___  ___  ________      
|\  \|\   ___  \|\___   ___\\  \|\  \|\   ___  \        Internal Network Tunnel  
\ \  \ \  \\ \  \|___ \  \_\ \  \\\  \ \  \\ \  \       
 \ \  \ \  \\ \  \   \ \  \ \ \  \\\  \ \  \\ \  \      Version: 1.0
  \ \  \ \  \\ \  \   \ \  \ \ \  \\\  \ \  \\ \  \     
   \ \__\ \__\\ \__\   \ \__\ \ \_______\ \__\\ \__\
    \|__|\|__| \|__|    \|__|  \|_______|\|__| \|__|      `

	return l
}

func main() {
	commandLine := flag.NewFlagSet("intun", flag.ExitOnError)
	i := commandLine.String("i", "", "--identity          identity of use (master, slave, client), required  ")
	h := commandLine.String("h", "", "--host=HOSTNAME     connect or start server host, required ")
	p := commandLine.Int("p", 0, "--port=PORT         connect or start server port, required ")
	commandLine.Parse(os.Args[1:])

	if *i == "" || *h == "" || *p == 0 {
		return
	}

	fmt.Println(logo())

	switch *i {
	case "master":
		master.LaunchMaster(*h, *p)
	case "slave":
	case "client":
		client.Launch(*h, *p)
	default:
		fmt.Printf("intun: identity (%s) is failed, use (master, slave, client)? \n", *i)
		return
	}

}
