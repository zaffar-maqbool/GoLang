package main

import "fmt"

func double(number int) int {
	return number * 2
}

// func doubleNumbers(numbers *[]int) []int {
// 	dNumbers := []int{}

// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, double(val))
// 	}
// 	return dNumbers
// }

func triple(number int) int {
	return number * 3
}

// func tripleNumbers(numbers *[]int) []int {
// 	dNumbers := []int{}

// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, triple(val))
// 	}
// 	return dNumbers
// }
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	transNumbers := []int{}

	for _, val := range *numbers {
		transNumbers = append(transNumbers, transform(val))
	}
	return transNumbers
}
func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)

	fmt.Println(doubled)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)
}
