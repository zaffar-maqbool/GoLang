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
	userAge, _ := strconv.ParseInt(userInput, 0, 64)

	if userAge >= 18 {
		fmt.Println("Welcome to club")
	} else {
		fmt.Println("Sorry YU Cant join")
	}

}
