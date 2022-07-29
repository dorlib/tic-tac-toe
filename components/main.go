package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello ! welcome to Tic Tac Toe Game")

	player1, player2 := setPlayers()
	tide := 0

	gameBoard := initBoard()

	symbols := gameBoard

	var point *Score

	point.player1Score = 0
	point.player2Score = 0
	point.tide = 0

	printBoard(symbols, player1, player2, tide)

	var gameOver bool
	gameOver = false
	newGame(player1, player2)

}
