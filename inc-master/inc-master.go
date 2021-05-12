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
 \ \  \ \  \\ \  \   \ \  \ \ \  \\\  \ \  \\ \  \      ROOT
  \ \  \ \  \\ \  \   \ \  \ \ \  \\\  \ \  \\ \  \     
   \ \__\ \__\\ \__\   \ \__\ \ \_______\ \__\\ \__\
    \|__|\|__| \|__|    \|__|  \|_______|\|__| \|__|      
`

	return l
}

func main() {
	fmt.Println(logo())

	commandLine := flag.NewFlagSet("inc", flag.ExitOnError)
	h := commandLine.String("h", "", "--host=HOSTNAME     start server host, required ")
	p := commandLine.Int("p", 0, "--port=PORT         start server port, required ")
	pw := commandLine.String("pw", "", "--password     for client auth, optional ")
	commandLine.Parse(os.Args[1:])

	if *h == "" || *p == 0 {
		return
	}

	address := fmt.Sprintf("%s:%d", *h, *p)
	_ = inc.LaunchRoot(address, *pw)

	fmt.Println("launch on", address)
	select {}

}
