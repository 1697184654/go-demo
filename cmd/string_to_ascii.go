package main

import "fmt"

func main() {
	var str string
	fmt.Println("请输入需要转换ascii码的字符串：")
	fmt.Scanln(&str)
	for _, char := range []rune(str) {
		fmt.Println(char)
	}
}
