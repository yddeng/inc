package main

import (
	"fmt"
	"github.com/yddeng/dutil/strutil"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func readLine1() string {
	buffer := make([]byte, 128)
	n, _ := os.Stdin.Read(buffer)
	return string(buffer[:n])
}

func readLine() string {
	buffer := make([]byte, 128)
	n, _ := os.Stdin.Read(buffer)
	return string(buffer[:n-1])
}

func readWords() (string, []string, int) {
	cmd := readLine()
	words := strutil.Str2Slice(cmd)
	wordsLen := len(words)
	return cmd, words, wordsLen
}

func command(name string, argv []string) *exec.Cmd {
	cmd := exec.Command(name, argv[1:]...)

	f, _ := os.Create("log.txt")
	w := io.MultiWriter(os.Stdout, f)
	cmd.Stdin = os.Stdin
	cmd.Stdout = w
	cmd.Stderr = os.Stderrssh
	//stdin, _ := cmd.StdinPipe()
	return cmd
}

func main() {
	signal.Ignore(syscall.SIGINT, syscall.SIGTERM)

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
	default:
		cmd := command(words[0], words)
		cmd.Run()
		goto loop
	}
}
