package main

import (
	"fmt"
	"math/rand"
)

// input as INT     output INT
func Add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

// naming a return type
func nameReturnValue(n1 int, n2 int) (sum int) {
	sum = n1 + n2
	return
}

// Function that return nothing
func newLine(number int) {
	fmt.Println("The Number Is :", number)
}

// Returning Multiple Values
func generateRandomNums() (int, int) {
	randNum1 := rand.Intn(100)
	randNum2 := rand.Intn(100)
	return randNum1, randNum2
}

func main() {
	//a := 5
	//b := 10
	//sum := a + b
	a, b := generateRandomNums()
	sum := Add(a, b)
	//fmt.Println(sum)
	newLine(sum)
	//fmt.Println(generateRandomNums())
	//fmt.Println(Add(generateRandomNums()))
	fmt.Print(nameReturnValue(4, 5))
}
