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
	userAge, _ := strconv.ParseInt(userAgeinput, 0, 64)

	if userAge >= 18 {
		fmt.Println("Welcome to club")
	}

}
