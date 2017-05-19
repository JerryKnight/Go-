package main

import "fmt"

func test(x int) {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer func() {
		fmt.Println(100 / x)
	}()
	defer fmt.Println("c")
	defer fmt.Println("d")
}

func main() {
	test(0)
}
