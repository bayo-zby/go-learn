package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 将值赋值为一个函数
	fn := func() string {
		return "function called"
	}

	fmt.Println(reflect.TypeOf(fn))
	fn()

	// fn()表示运行函数返回结果
	// fn则指向函数本身
	fmt.Println(typetest(fn))

	// 传递函数时，不能使用已创建的函数
	// tt := typetest()
	// fmt.Println(reflect.TypeOf(tt))

}

func typetest(f func() string) string {
	return f()
}
