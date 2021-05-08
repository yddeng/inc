package main

import (
	"fmt"
	"github.com/yddeng/dutil/strutil"
	"os"
	"os/exec"
	"syscall"
)

func readLine() string {
	buffer := make([]byte, 128)
	n, _ := os.Stdin.Read(buffer)
	return string(buffer[:n-1])
}

func readWords() (string, []string, int) {
	cmd := readLine()
	words := strutil.Str2Slice(cmd, " ")
	wordsLen := len(words)
	return cmd, words, wordsLen
}

func replaceExec(name string, argv []string) error {
	binary, err := exec.LookPath(name)
	if err != nil {
		return err
	}

	err = syscall.Exec(binary, argv, []string{})
	if err != nil {
		return err
	}
	return nil
}

func main() {

	fmt.Println("exec")

	_, words, _ := readWords()

	if err := replaceExec(words[0], words); err != nil {
		fmt.Println(err)
	}

	fmt.Println("exec end")

}
