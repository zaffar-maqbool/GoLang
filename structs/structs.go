package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func getUserData(promtText string) string {
	fmt.Println(promtText)
	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.Replace(userInput, "\n", " ", -1)
	return cleanedInput
}

func main() {
	firstName := getUserData("Please Enter Your First Name: ")
	lastName := getUserData("Please Enter Your Last Name: ")
	birthdate := getUserData("Please Enter Your Birthdate in (MM/DD/YYYY: ")
	created := time.Now()

	fmt.Println(firstName, lastName, birthdate, created)

}
