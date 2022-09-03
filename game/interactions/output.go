package interactions

import (
	"fmt"
)

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting new game ...")
	fmt.Println("Good Luck:")
}
func ShowAvailableActions(isSpecialAttackAvailable bool) {
	fmt.Println("Please Chose Your Action")
	fmt.Println("1: Attack Monster")
	fmt.Println("2: Heal")

	if isSpecialAttackAvailable {
		fmt.Println("3: Special Attack ")
	}
}
