package main

import (
	"flag"
	"fmt"
	"github.com/yddeng/inc"
	"github.com/yddeng/inc/net"
	net2 "net"
	"os"
	"strconv"
)

func main() {
	fmt.Println(inc.Logo("inc-slave"))

	commandLine := flag.NewFlagSet("inc", flag.ExitOnError)
	a := commandLine.String("a", "", "--address     address of external master, required ")
	n := commandLine.String("n", "slave", "--name        name for slave, optional ")
	r := commandLine.String("r", "", "--register    register (internalIP, internalPort, externalPort, description)")
	commandLine.Parse(os.Args[1:])

	if *a == "" {
		return
	}

	_ = inc.LaunchIncSlave(*n, *a, parseRe(*r))

	select {}
}

func parseRe(s string) *net.Mapping {
	if s == "" {
		return nil
	}

	workds, length := inc.ReadWords(s)
	if length != 4 {
		fmt.Println("register str failed:", s)
		os.Exit(0)
	}

	ip := net2.ParseIP(workds[0])
	if ip == nil {
		fmt.Println("register internal ip failed:", workds[0])
		os.Exit(0)
	}
	inPort, err := strconv.Atoi(workds[1])
	if err != nil {
		fmt.Println("register internal port failed:", workds[1])
		os.Exit(0)
	}
	exPort, err := strconv.Atoi(workds[2])
	if err != nil {
		fmt.Println("register external port failed:", workds[2])
		os.Exit(0)
	}

	return &net.Mapping{
		InternalIp:   workds[0],
		InternalPort: int32(inPort),
		ExternalPort: int32(exPort),
		Description:  workds[3],
	}
}
