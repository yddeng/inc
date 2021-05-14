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

func Usage() {
	usage := `Usage of inc-slave:
  inc-slave [Options]

General options:
  -n, --name         name for slave (default: "slave"). 
  -r, --register     register agents, use ';' split each other (e.g '127.0.0.1 22 2201 ssh; 127.0.0.1 5432 2345 postgreSQL').

Connection options:
  -h, --host         master server host.
  -p, --port         master server port (default: "9852").`

	fmt.Println(usage)
}

func main() {
	commandLine := flag.NewFlagSet("inc-slave", flag.ExitOnError)
	commandLine.Usage = Usage
	h := commandLine.String("h", "", "--host     master server host, required ")
	p := commandLine.Int("p", 9852, "--port     master server port, required ")
	n := commandLine.String("n", "slave", "--name        name for slave, optional. ")
	r := commandLine.String("r", "", "--register    register agents, use ';' split each other (e.g '127.0.0.1 22 2201 ssh; 127.0.0.1 5432 2345 postgreSQL').")
	commandLine.Parse(os.Args[1:])

	fmt.Println(inc.Logo("inc-slave"))

	var mappings []*net.Mapping
	if *r != "" {
		ss := strings.Split(*r, ";")
		mappings = make([]*net.Mapping, 0, len(ss))
		for _, s := range ss {
			if m := parseMapping(s); m != nil {
				mappings = append(mappings, m)
			}
		}
	}

	address := fmt.Sprintf("%s:%d", *h, *p)
	_ = inc.LaunchIncSlave(*n, address, mappings)

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
