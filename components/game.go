package main

import (
	"fmt"
	"math/rand"
)

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

	// for loop for every round of the game
	for i := 0; i < 9 || gameOver == false; i++ {
		if i%2 == 0 {
			fmt.Printf("Next player to make a mark is %v", starter.name)
			printBoard(starter, seconder, tide)
			turnIndex++
		} else {
			fmt.Printf("Next player to make a mark is %v", seconder.name)
			turnIndex++
		}
	}
}
