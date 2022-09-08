package main

func cons() {

	const userName = "ZAFFAR"
	const age = 44 / 2

	// const random = rand.Int() // not allowed

	const (
		// if order is in increamental like 1 2 3 4 ...
		// we can use "iota"
		inputAttack = iota + 1 // = 0 + 1
		inputSpecialAttack
		inputHeal

		dummy = iota + 11 // 3 + 11
		incre             // 15

	)
}
