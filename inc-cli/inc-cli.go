package main

import (
	"flag"
	"fmt"
	"github.com/yddeng/inc"
	"github.com/yddeng/utils/strutil"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
)

func logo() string {
	l := `
___  ________   _________  ___  ___  ________      
|\  \|\   ___  \|\___   ___\\  \|\  \|\   ___  \        Internal Network Tunnel  
\ \  \ \  \\ \  \|___ \  \_\ \  \\\  \ \  \\ \  \       
 \ \  \ \  \\ \  \   \ \  \ \ \  \\\  \ \  \\ \  \      Client
  \ \  \ \  \\ \  \   \ \  \ \ \  \\\  \ \  \\ \  \     
   \ \__\ \__\\ \__\   \ \__\ \ \_______\ \__\\ \__\
    \|__|\|__| \|__|    \|__|  \|_______|\|__| \|__|      
`

	return l
}

var buffer = make([]byte, 128)

func readLine() string {
	n, _ := os.Stdin.Read(buffer)
	return string(buffer[:n-1])
}

func readWords() (string, []string, int) {
	line := readLine()
	words := strutil.Str2Slice(line)
	wordsLen := len(words)
	return line, words, wordsLen
}

func main() {
	signal.Ignore(syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(logo())

	commandLine := flag.NewFlagSet("inc", flag.ExitOnError)
	a := commandLine.String("a", "", "--host=HOSTNAME     start server host, required ")
	commandLine.Parse(os.Args[1:])

	if *a == "" {
		return
	}

	fmt.Print("Password to connection master:")
	pw := readLine()

	address := fmt.Sprintf("%s:%d", *h, *p)
	client := inc.LaunchClient(address, pw)

loop:
	fmt.Print("==>")
	_, words, length := readWords()
	switch length {
	case 0:
		goto loop
	case 1:
		switch words[0] {
		case "quit":
			return
		default:
			goto loop
		}
	case 2:
		switch words[0] {
		case "select":
			num, err := strconv.Atoi(words[1])
			if err != nil {
				fmt.Println(err)
				goto loop
			}
			client.SelectLeaf(uint32(num))
			goto loop
		default:
			goto loop
		}
	case 3:
		switch words[0] {
		case "create":
			if err := client.CreateTunnel(words[1], words[2]); err != nil {
				fmt.Println(err)
			}

			goto loop
		default:
			goto loop
		}
	default:
		cmd := cmdExec(words[0], words)
		cmd.Run()
		goto loop
	}

}

func register() {

}
