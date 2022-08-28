package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("This is a basic BMI calculator")

	fmt.Print("Please enter yours weight in KG: ")
	weightInput, _ := reader.ReadString('\n')
	fmt.Print("Please Enter Your Hight in M ")
	heightInput, _ := reader.ReadString('\n')

	// fmt.Print(weightInput)
	// fmt.Print(heightInput)

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)
	// converting next line

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)

	fmt.Printf("your BMI : %.2f", bmi)
}
