package main

import "math/rand"

func newGame(player1 *Player, player2 *Player) {
	var starter *Player
	var seconder *Player

	//toss a coin to determine who's starts
	result := rand.Intn(2)
	if result == 0 {
		starter = player1
		seconder = player2
	} else {
		starter = player2
		seconder = player1
	}

	turnIndex := 0

	starterTurns = []int{0, 2, 4, 6, 8}
	seconderTurns = []int{1, 3, 5, 7}

	// for loop for every round of the game
	for i := 0; i < 9 || gameOver == false; i++ {

	}
}
