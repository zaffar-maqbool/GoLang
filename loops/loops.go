package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserChoice() {
	fmt.Println("please make yur choice")
	fmt.Println("1) Add upp all the numbers of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) sum uo manually entered numbers")
	fmt.Println("4) sum up a list of entered numbers")

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n ", " ", -1)
}
func main() {

}
