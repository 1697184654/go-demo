package main

import (
	"fmt"
	"os"
)

// go run cmd/args_command.go param1 param2 param3
func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
}
