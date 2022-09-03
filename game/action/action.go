package action

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())

var randGenerate = rand.New(randSource)

var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttackAvailable bool) (int, int) {
	minAttack := PLAYER_ATTACK_MIN_DMG
	maxAttack := PLAYER_ATTACK_MIN_DMG

	if isSpecialAttackAvailable {
		minAttack = PLAYER_ATTACK_MIN_DMG
		maxAttack = PLAYER_ATTACK_MAX_DMG
	}
	dmgValue := generateRandBetween(minAttack, maxAttack)

	//currentMonsterHealth = currentMonsterHealth - dmgValue
	currentMonsterHealth -= dmgValue
	return currentMonsterHealth, currentPlayerHealth
}

func AttackPlayer() (int, int) {
	minAttack := MONSTER_ATTACK_MIN_DMG
	maxAttack := MONSTER_ATTACK_MAX_DMG

	dmgValue := generateRandBetween(minAttack, maxAttack)

	//currentMonsterHealth = currentMonsterHealth - dmgValue
	currentPlayerHealth -= dmgValue
	return currentMonsterHealth, currentPlayerHealth
}

func HealPlayer() (int, int) {

	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)
	healDiff := 100 - currentPlayerHealth
	if healDiff >= healValue {
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
	}
	currentPlayerHealth += healValue
	return currentMonsterHealth, currentPlayerHealth
}

func generateRandBetween(min int, max int) int {
	return randGenerate.Intn(max-min) + min
}
