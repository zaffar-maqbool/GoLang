package main

import (
	"fmt"

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

	fmt.Println("")
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
