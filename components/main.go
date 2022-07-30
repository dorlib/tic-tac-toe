package main

import (
	"fmt"
)

func main() {
	fmt.Println("###########################################")
	fmt.Println("### Hello ! welcome to Tic Tac Toe Game ###")
	fmt.Println("###########################################")

	player1, player2 := setPlayers()
	var tide int
	var gameOver bool
	var gotTide bool
	var gameBoard [][]string

	for {
		gameBoard = initBoard()
		gameOver, gotTide = newGame(player1, player2, tide, gameBoard)
		if gameOver {
			if gotTide {
				tide++
			}
			fmt.Println("game over")
			break
		}
	}
}
