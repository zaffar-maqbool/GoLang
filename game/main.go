package main

import (
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

	if userChoice == "1" {
		return "ATTACK"
	} else if userChoice == "2" {
		return "HEAL"
	} else {

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
