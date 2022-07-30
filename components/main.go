package main

import (
	"fmt"
)

func main() {
	fmt.Println("###########################################")
	fmt.Println("### Hello ! welcome to Tic Tac Toe Game ###")
	fmt.Println("###########################################")

	player1, player2 := setPlayers()

	gameBoard := initBoard()

	var points *Score

	var gameOver bool
	gameOver = newGame(player1, player2, points, gameBoard)
	if gameOver {
		fmt.Println("game over")
	}
}
