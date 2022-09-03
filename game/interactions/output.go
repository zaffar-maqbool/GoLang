package interactions

import (
	"fmt"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

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

func DeclareWinner(winner string) {
	fmt.Println("GAME OVER")
	fmt.Printf("%v Won!\n", winner)
}

func PrintRoundStatstics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player Attacked Monster for %v damage \n", roundData.PlayerHealValue)

	} else if roundData.Action == "SPECIAL ATTACK" {
		fmt.Printf("Player Attacked Monster With Special Attack for %v damage", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player Healed for  %v \n", roundData.PlayerAttackDmg)
	}
	fmt.Printf(" Monster Attacked  Player for %v damage", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health %c is ", roundData.PlayerHealth)
	fmt.Printf("Monster Health %c is ", roundData.MonsterHealth)
}
