package action

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())

var randGenerate = rand.New(randSource)

var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttackAvailable bool) {
	minAttack := PLAYER_ATTACK_MIN_DMG
	maxAttack := PLAYER_ATTACK_MIN_DMG

	if isSpecialAttackAvailable {
		minAttack = 10
		maxAttack = 20
	}
	dmgValue := generateRandBetween(minAttack, maxAttack)

	//currentMonsterHealth = currentMonsterHealth - dmgValue
	currentMonsterHealth -= dmgValue

}

func AttackPlayer() {
	minAttack := 9
	maxAttack := 13

	dmgValue := generateRandBetween(minAttack, maxAttack)

	//currentMonsterHealth = currentMonsterHealth - dmgValue
	currentPlayerHealth -= dmgValue
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

func generateRandBetween(min int, max int) int {
	return randGenerate.Intn(max-min) + min
}
