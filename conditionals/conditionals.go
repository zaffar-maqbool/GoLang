package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your age ")
	userInput, _ := reader.ReadString('\n')
	userAge, err := strconv.ParseInt(userInput, 0, 64)
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
