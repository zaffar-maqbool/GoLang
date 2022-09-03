package action

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())

var randGenerate = rand.New(randSource)

var currentMonsterHealth = 100
var currentPlayersHealth = 100

func AttackMonster(isSpecialAttackAvailable bool) {
	minAttack := 5
	maxAttack := 10

	if isSpecialAttackAvailable {
		minAttack = 10
		maxAttack = 20
	}
	dmgValue := generateRandBetween(minAttack, maxAttack)
	currentMonsterHealth = currentMonsterHealth - dmgValue
	return
}

func HealPlayer() {

}

func AttackPlayer() {

}
func generateRandBetween(min int, max int) int {
	return randGenerate.Intn(max-min) + min
}
