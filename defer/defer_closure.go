package main

import "fmt"

func test() {
	x, y := 10, 20

	defer func(i int) {
		fmt.Println("i:", i)
		fmt.Println("y:", y)
	}(x)

	x += 10
	y += 100
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}

func main() {
	test()
}
