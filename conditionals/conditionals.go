package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserAge() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input your age ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	userAge, err := strconv.ParseInt(userInput, 0, 64)

	return int(userAge), err
}

func main() {
	userAge, err := getUserAge()
	fmt.Println(err)
	if err != nil {
		fmt.Println("invalid input")
		return
	}
	//userAge := 30
	//isOldEnough := userAge >= 18
	if userAge >= 30 && userAge < 50 || userAge >= 60 {
		fmt.Println("Welcome to club")
	} else if userAge >= 18 {
		fmt.Print("Welcome")
	} else if userAge > 50 {
		fmt.Println("old")
	} else {
		fmt.Println("Sorry YU Cant join")
	}

}
