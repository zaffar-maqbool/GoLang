package main

import "fmt"

func double(number int) int {
	return number * 2

}

func main() {
	age := 22
	fmt.Println(age)

	myAge := &age
	fmt.Println(myAge)
	fmt.Println(*myAge)

	*myAge = 33
	fmt.Println(age)
	doubleAge := double(age)
	fmt.Println(doubleAge)
}
