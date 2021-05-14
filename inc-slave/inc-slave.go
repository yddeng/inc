package main

import (
	"flag"
	"fmt"
	"github.com/yddeng/inc"
	"github.com/yddeng/inc/net"
	net2 "net"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(inc.Logo("inc-slave"))

	commandLine := flag.NewFlagSet("inc-slave", flag.ExitOnError)
	a := commandLine.String("a", "", "--address     address of external master, required. ")
	n := commandLine.String("n", "", "--name        name for slave, optional. ")
	r := commandLine.String("r", "", "--register    register one agent, e.g '127.0.0.1 22 2201 ssh'.")
	rs := commandLine.String("rs", "", "--registers    register multiple agents, use ';' split each other.")
	commandLine.Parse(os.Args[1:])

	if *a == "" {
		return
	}

	mappings := make([]*net.Mapping, 0, 4)
	if m := parseMapping(*r); m != nil {
		mappings = append(mappings, m)
	}

	if *rs != "" {
		ss := strings.Split(*rs, ";")
		for _, s := range ss {
			if m := parseMapping(s); m != nil {
				mappings = append(mappings, m)
			}
		}
	}

	_ = inc.LaunchIncSlave(*n, *a, mappings)

	select {}
}

func parseMapping(s string) *net.Mapping {
	if s == "" {
		return nil
	}

	worlds, length := inc.ReadWords(s)
	if length != 4 {
		fmt.Println("register command failed:", s)
		os.Exit(0)
	}

	ip := net2.ParseIP(worlds[0])
	if ip == nil {
		fmt.Printf("internal ip failed: %s(%s)", s, worlds[0])
		os.Exit(0)
	}
	inPort, err := strconv.Atoi(worlds[1])
	if err != nil {
		fmt.Printf("internal port failed: %s(%s)", s, worlds[1])
		os.Exit(0)
	}
	exPort, err := strconv.Atoi(worlds[2])
	if err != nil {
		fmt.Printf("external port failed: %s(%s)", s, worlds[2])
		os.Exit(0)
	}

	return &net.Mapping{
		InternalIp:   worlds[0],
		InternalPort: int32(inPort),
		ExternalPort: int32(exPort),
		Description:  worlds[3],
	}
}
