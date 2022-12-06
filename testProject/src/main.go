package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello go")
	s := "golang"
	println(s)
	var a int = 10
	var b int = 12
	var c int
	c = GetVal(a, b)
	fmt.Println(c)
}
