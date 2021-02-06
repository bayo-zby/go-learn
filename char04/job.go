package main

import "fmt"

func main() {
	wenduzhuanhuan(10)
	recursion(1)
}

func wenduzhuanhuan(hs int) int {
	// 华氏转摄氏
	return 0
}

func recursion(num int) int {
	fmt.Println(num)
	if num >= 10 {
		return num
	}
	num++
	return recursion(num)
}

func test01(a, b int) (c, d, e int) {
	return
}
