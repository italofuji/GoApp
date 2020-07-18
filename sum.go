package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Sum result: ", sum(5, 5))
}