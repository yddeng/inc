package main

import (
	"flag"
	"fmt"
	"github.com/yddeng/inc"
	"os"
)

func main() {
	fmt.Println(inc.Logo("inc-master"))

	commandLine := flag.NewFlagSet("inc", flag.ExitOnError)
	h := commandLine.String("h", "", "--host=HOSTNAME     start server host, required ")
	p := commandLine.Int("p", 0, "--port=PORT         start server port, required ")
	t := commandLine.String("t", "", "--token     for client auth, optional ")
	commandLine.Parse(os.Args[1:])

	if *h == "" || *p == 0 {
		return
	}

	address := fmt.Sprintf("%s:%d", *h, *p)
	_ = inc.LaunchIncMaster(*h, *p, *t)

	fmt.Println("launch on", address)
	select {}

}
