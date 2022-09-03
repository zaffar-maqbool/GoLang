package action

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())

var randGenerate = rand.New(randSource)

var currentMonsterHealth = 100
var currentPlayerHealth = 100

func AttackMonster(isSpecialAttackAvailable bool) {
	minAttack := 5
	maxAttack := 10

	if isSpecialAttackAvailable {
		minAttack = 10
		maxAttack = 20
	}
	dmgValue := generateRandBetween(minAttack, maxAttack)

	//currentMonsterHealth = currentMonsterHealth - dmgValue
	currentMonsterHealth -= dmgValue

}

func HealPlayer() {
	minHeal := 10
	maxHeal := 20

	healValue := generateRandBetween(minHeal, maxHeal)
	healDiff := 100 - currentPlayerHealth
	if healDiff >= healValue {
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = 100
	}
	currentPlayerHealth += healValue
}

func AttackPlayer() {
	minAttack := 9
	maxAttack := 13

	dmgValue := generateRandBetween(minAttack, maxAttack)

	//currentMonsterHealth = currentMonsterHealth - dmgValue
	currentPlayerHealth -= dmgValue
}
func generateRandBetween(min int, max int) int {
	return randGenerate.Intn(max-min) + min
}
