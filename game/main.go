package main

func startGme() {

}
func executeRound() string {

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
