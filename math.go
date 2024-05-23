package main

import "fmt"

func main() {
	fmt.Println(add(10, 10))
}

func add(a int, b int) int {
	return a + b
}
