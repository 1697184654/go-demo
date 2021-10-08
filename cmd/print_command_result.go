package main

import (
	"fmt"
	"os"
	"os/exec"
)

// go run cmd/print_command-result.go ps aux
func main() {
	str, err := runCommand(os.Args[1:]...)
	if err != nil {
		fmt.Println("Command Failed:", err.Error())
	}
	fmt.Println(string(str))
}

func runCommand(arg ...string) ([]byte, error) {
	if len(arg) == 0 {
		fmt.Println("Usage: %s args...\n", os.Args[0])
		os.Exit(-1)
	}
	return exec.Command(arg[0], arg[1:]...).Output()
}
