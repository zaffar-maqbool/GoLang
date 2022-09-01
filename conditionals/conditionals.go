package main

import (
	"fmt"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Please input your age ")
	// userInput, _ := reader.ReadString('\n')
	// userAge, _ := strconv.ParseInt(userInput, 0, 64)

	userAge := 30
	//isOldEnough := userAge >= 18
	if userAge >= 30 {
		fmt.Println("Welcome to club")
	} else if userAge >= 18 {
		fmt.Print("Welcome")
	} else {
		fmt.Println("Sorry YU Cant join")
	}

}
