package main

import (
	"github.com/zaffar-maqbool/Golang/game/action"
	"github.com/zaffar-maqbool/Golang/game/interactions"
)

var currentRound = 0

func startGme() {
	interactions.PrintGreeting()
}
func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0
	//to know wheather speacial attack can be used
	interactions.ShowAvailableActions(isSpecialRound)
	userChoice := interactions.GetPlayerChoice(isSpecialRound)

	var playerHealth int
	var monsterHealth int
	if userChoice == "1" {
		action.AttackMonster(false)
	} else if userChoice == "2" {
		action.HealPlayer()
	} else {
		action.AttackMonster(true)

	}
	action.AttackPlayer()
	playerHealth, monsterHealth = action.GetHealth()

	if playerHealth <= 0 {
		return "Monster"

	} else if monsterHealth <= 0 {
		return "Player"
	}
	return ""

}
func endGame() {

}

func main() {

	startGme()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}
