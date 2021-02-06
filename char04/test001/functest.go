package main

import "fmt"

func sayHi() (x, y string) {
	fmt.Printf("%q %q", x, y)
	return
}

func main() {
	sayHi()
}
