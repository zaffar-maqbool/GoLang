package main

import "fmt"

func main() {
	factorialNum := factorial(5)
	fmt.Println(factorialNum)
}
func factorial(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}
