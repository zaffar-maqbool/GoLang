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
	//hobbies := []string{"Cook", "Code"}
	hobbies := make([]string, 2, 10)
	hobbies[1] = "CODE"
	hobbies[0] = "COOK"
	hobbies = append(hobbies, "LEARN", "TEACH")
	fmt.Println(hobbies)

	// aMap := make(map[string]int, 2)

	moreHobbies := new([]string)
	fmt.Println(moreHobbies)
	*moreHobbies = append(*moreHobbies, "WORK")
	fmt.Println(*moreHobbies)

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
