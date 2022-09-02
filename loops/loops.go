package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserChoice() (string, error) {
	fmt.Println("please make yur choice")
	fmt.Println("1) Add upp all the numbers of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) sum uo manually entered numbers")
	fmt.Println("4) sum up a list of entered numbers")

	fmt.Print("Please make your choice: ")
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n ", " ", -1)
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
	chosenNum, err := getInputNumber()
	if err != nil {
		fmt.Println("Invalid number input")
		return
	}
	fmt.Println(chosenNum)
	//                 i = i + 1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}
func calFactorial() {

}
func calSumManually() {

}
func calListSum() {

}
func getInputNumber() (int, error) {
	inputNum, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	inputNum = strings.Replace(inputNum, "\r\n", "", -1)
	chosenNum, err := strconv.ParseInt(inputNum, 0, 64)
	if err != nil {

		return 0, err
	}
	return int(chosenNum), nil
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
