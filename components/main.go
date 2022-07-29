package main

import (
	"fmt"
)

//type Player struct {
//	sign  string
//	name  string
//	score int
//}
//
//
//func newBoard(square [][]Symbol, dimention int) *Board {
//	b := Board{}
//}

func main() {
	fmt.Println("Hello ! welcome to Tic Tac Toe Game")

	player1, player2 := setPlayers()
	tide := 0

	gameBoard := initBoard()

	symbols := gameBoard
	printBoard(symbols, player1, player2, tide)

	var gameOver bool
	gameOver = false

}
