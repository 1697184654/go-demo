package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Match("第1集 《小蝌蚪找妈妈》：识字记词", "1 小蝌蚪找妈妈"))
	fmt.Println(Match("第2集 《小蝌蚪找妈妈》课文赏析", "1 小蝌蚪找妈妈"))
	fmt.Println(Match("第3集 《我是什么》：识字记词", "2 我是什么"))
	fmt.Println(Match("第4集 《我是什么》课文赏析", "2 我是什么"))
	fmt.Println(Match("s2 我是什么f", "b第4集 《我是什么》课文赏析a"))
	fmt.Println(Match("s2 我是什么f", "第1集 《小蝌蚪找妈妈》：识字记词"))
}

func Match(str1 string, str2 string) float64 {
	var char1 = []byte(str1)
	var char2 = []byte(str2)
	var len1 int = len(str1)
	var len2 int = len(str2)
	var dif [200][200]int
	for i := 0; i <= len1; i++ {
		dif[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dif[0][j] = j
	}
	var temp int
	for m := 1; m <= len1; m++ {
		for n := 1; n <= len2; n++ {
			if char1[m-1] == char2[n-1] {
				temp = 0
			} else {
				temp = 1
			}
			dif[m][n] = int(math.Min(float64(dif[m-1][n-1]+temp), float64(dif[m][n-1]+1)))
			dif[m][n] = int(math.Min(float64(dif[m][n]), float64(dif[m-1][n]+1)))
		}
	}
	x := 1 - float64(float64(dif[len1][len2])/float64(math.Max(float64(len1), float64(len2))))
	return x
}
