package main

import "fmt"

func main() {
	//hobbies := []string{"Cook", "Code"}
	hobbies := make([]string, 2, 10)
	hobbies[1] = "CODE"
	hobbies[0] = "COOK"
	hobbies = append(hobbies, "LEARN", "TEACH")
	fmt.Println(hobbies)

	// aMap := make(map[string]int, 2)

	moreHobbies := new([]string)
	fmt.Println(moreHobbies)
	*moreHobbies = append(*moreHobbies, "WORK")
	fmt.Println(*moreHobbies)

}
