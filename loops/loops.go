package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserChoice() (string, error) {
	fmt.Println("Please make your choice")
	fmt.Println("1) Add up all the numbers of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered numbers")
	fmt.Println("4) Sum up a list of entered numbers")

	fmt.Print("Please make your choice: ")
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}

	if userInput == "1" ||
		userInput == "2" ||
		userInput == "3" ||
		userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("INVALID INPUT")
	}
}
func calSumUpToNumber() {
	fmt.Println("Plese enter your number: ")
	chosenNum, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid number input")
		return
	}

	fmt.Println(chosenNum)
	sum := 0
	//                 i = i + 1
	for i := 1; i <= chosenNum; i++ {
		sum = sum + i
	}
	fmt.Printf("Result is %v ", sum)

}
func calFactorial() {

}
func calSumManually() {

}
func calListSum() {

}

func getInputNumber() (int, error) {
	inputNumber, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	if runtime.GOOS == "windows" {
		inputNumber = strings.Replace(inputNumber, "\r\n", "", -1)
	} else {
		inputNumber = strings.Replace(inputNumber, "\n", "", -1)
	}
	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)

	if err != nil {
		return 0, err
	}

	return int(chosenNumber), nil
}

func main() {
	selectedChoice, err := getUserChoice()

	if err != nil {
		fmt.Println("invalid choice , exiting ")
		return
	}
	if selectedChoice == "1" {
		calSumUpToNumber()
	} else if selectedChoice == "2" {
		calFactorial()
	} else if selectedChoice == "4" {
		calSumManually()
	} else {
		calListSum()
	}

}
