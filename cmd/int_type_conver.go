package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int64
	fmt.Println("类型转换：")
	fmt.Scanln(&x)

	fmt.Println("二进制：", strconv.FormatInt(x, 2))
	fmt.Println("八进制：", strconv.FormatInt(x, 8))
	fmt.Println("十六进制：", strconv.FormatInt(x, 16))

}
