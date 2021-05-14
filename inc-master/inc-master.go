package main

import (
	"flag"
	"fmt"
	"github.com/yddeng/inc"
	"os"
)

func Usage() {
	usage := `Usage of inc-master:
  inc-master [Options]

Connection options:
  -h, --host         start server host.
  -p, --port         start server port (default: "9852").
  -t, --token        auth for client connection, optional.`

	fmt.Println(usage)
}

func main() {
	commandLine := flag.NewFlagSet("inc", flag.ExitOnError)
	commandLine.Usage = Usage
	h := commandLine.String("h", "", "--host     start server host, required ")
	p := commandLine.Int("p", 9852, "--port     start server port, required ")
	t := commandLine.String("t", "", "--token    for client auth, optional ")
	commandLine.Parse(os.Args[1:])

	fmt.Println(inc.Logo("inc-master"))

	if *h == "" || *p == 0 {
		return
	}

	address := fmt.Sprintf("%s:%d", *h, *p)
	_ = inc.LaunchIncMaster(*h, *p, *t)

	fmt.Println("launch on", address)
	select {}

}
