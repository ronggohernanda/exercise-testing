package main

import "fmt"

func Multiplication(a int, b int) int {
	return a * b
}

func Division(c int, d int) int {
	return c / d
}

func Modulus(x int, y int) int {
	return x % y
}

func main() {
	fmt.Println(Multiplication(5, 5))
	fmt.Println(Division(100, 5))
	fmt.Println(Modulus(10, 3))
}
