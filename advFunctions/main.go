package main

import "fmt"

func double(number int) int {
	return number * 2
}
func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, double(val))
	}
	return dNumbers
}

func triple(number int) int {
	return number * 3
}
func tripleNumbers(numbers *[]int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, triple(val))
	}
	return dNumbers
}

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumbers(&numbers)

	fmt.Println(doubled)
	tripled := tripleNumbers(&numbers)
	fmt.Println(tripled)
}
