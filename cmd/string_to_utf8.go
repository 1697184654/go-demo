package main

import "fmt"

func main() {
	var str string
	fmt.Println("golang 默认utf-8编码，可以编解码所有unicode字符集")
	fmt.Println("请输入需要转换utf-8编码的字符串：")
	fmt.Scanln(&str)
	for _, char := range []rune(str) {
		fmt.Println(char)
	}
}
