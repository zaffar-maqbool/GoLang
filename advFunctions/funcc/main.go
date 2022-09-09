package main

import "fmt"

type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

// package main

// import "fmt"

// type transformFn func(int) int

// func double(number int) int {
// 	return number * 2
// }

// // func doubleNumbers(numbers *[]int) []int {
// // 	dNumbers := []int{}

// // 	for _, val := range *numbers {
// // 		dNumbers = append(dNumbers, double(val))
// // 	}
// // 	return dNumbers
// // }

// func triple(number int) int {
// 	return number * 3
// }

// // func tripleNumbers(numbers *[]int) []int {
// // 	dNumbers := []int{}

// // 	for _, val := range *numbers {
// // 		dNumbers = append(dNumbers, triple(val))
// // 	}
// // 	return dNumbers
// // }
// func transformNumbers(numbers *[]int, transform transformFn) []int {
// 	transNumbers := []int{}

// 	for _, val := range *numbers {
// 		transNumbers = append(transNumbers, transform(val))
// 	}
// 	return transNumbers
// }
// func getTrnsFunc() (numbers *[]int) transformFn {
// 	if (*numbers) [0] == 1{
// 		return double
// 	}else {
// 		return triple
// 	}
// }
// func main() {
// 	numbers := []int{1, 2, 3, 4}
// 	moreNumbers := []{5,4,3}
// 	doubled := transformNumbers(&numbers, double)

// 	fmt.Println(doubled)
// 	tripled := transformNumbers(&numbers, triple)
// 	fmt.Println(tripled)
// 	trnsformerFn1 := getTrnsFunc(&numbers)
// 	trnsformerFn2 := getTrnsFunc(&moreNumbers)
// 	transformedNumber := transNumbers(&numbers, trnsformerFn1)
// 	moreTransformedNumbers := transNumbers(&moreNumbers,trnsformerFn2)
// 	fmt.Println(trnsformerFn1)
// 	fmt.Println(trnsformerFn2)
// }
