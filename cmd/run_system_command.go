package main

import (
	"fmt"
	"os/exec"
)

func main() {
	str, err := runCommand("ls -l")
	if err != nil {
		fmt.Println("Command Failed:", err.Error())
	}
	fmt.Println(string(str))
}

func runCommand(cmd string) ([]byte, error) {
	return exec.Command("/bin/sh", "-c", cmd).Output()
}
