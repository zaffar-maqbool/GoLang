package interactions

import (
	"fmt"
	"os"
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

func PrintRoundStatstics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player Attacked Monster for %v damage\n", roundData.PlayerAttackDmg)

	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player Attacked Monster With Special Attack for %v damage", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player Healed for  %v\n", roundData.PlayerHealValue)
	}
	fmt.Printf("Monster Attacked  Player for %v damage\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health %v is\n ", roundData.PlayerHealth)
	fmt.Printf("Monster Health %v is\n ", roundData.MonsterHealth)
}
func DeclareWinner(winner string) {
	fmt.Println("GAME OVER")
	fmt.Printf("%v Won!\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {
	file, err := os.Create("gameLog.txt")
	if err != nil {
		fmt.Println("Saving a log file failed.Exiting")
		return
	}
	for i, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(i + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health ":        fmt.Sprint(value.PlayerHealth),
			"Monster Health ":       fmt.Sprint(value.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing log file failed . Exiting ")
			continue
		}
	}
	file.Close()
	fmt.Println("Writing Log Data Successfull ")
}
