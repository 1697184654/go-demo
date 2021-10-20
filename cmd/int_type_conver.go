package main

import (
	"fmt"
)

func main() {
	var x int64
	fmt.Println("类型转换：")
	fmt.Scanln(&x)

	//strconv.FormatInt(x, 2)  整数转二进制
	fmt.Printf("二进制原码：%b \n", x)
	fmt.Printf("二进制反码：%b \n", 1^1)
	fmt.Printf("八进制：%o \n", x)
	fmt.Printf("十六进制：%x \n", x)

}
