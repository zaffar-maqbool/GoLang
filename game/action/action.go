package action

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())

var randGenerate = rand.New(randSource)

func AttackMonster(isSpecialAttackAvailable bool) {

}

func HealPlayer() {

}

func AttackPlayer() {

}
func generateRandBetween() {
	return randGenerate.Intn()
}
