package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	err := exec.Command("jemeter").Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
