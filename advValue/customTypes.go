package main

import "fmt"

type person struct {
	name string
	age  int
}

type customNumber int

type personData map[string]person

// func (number int) pow() {
// 	return number * 2
// }

func (number customNumber) pow(power int) customNumber {
	result := number

	for i := 1; i < power; i++ {
		result = result * number
	}
	return result
}

// main
func call() {

	var people []person = []person{
		{"Max", 32},
	}
	fmt.Println(people)
	var peps personData = personData{
		"p1": {"ZAFFAR", 22},
	}
	fmt.Println(peps)
	var dummyNumber customNumber = 5
	poweredNumber := dummyNumber.pow(3)
	fmt.Println(poweredNumber)
}
