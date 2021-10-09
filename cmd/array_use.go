package main

import "fmt"

func main() {
	SimpleArray()
}

func SimpleArray() {
	var x = [3]int{1, 2, 3}
	fmt.Println(x)
	fmt.Println(x[1:])
	var y = [...]int{1, 2, 3, 4}
	fmt.Println(y)
	fmt.Println(y[1:])
}
