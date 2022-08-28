package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zaffar-maqbool/Golang/calci/info"
)

func main() {
	fmt.Println(info.MainTitle)

	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')
	fmt.Print(info.HeightPrompt)
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
