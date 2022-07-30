package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// newGame responsible for the flow of a single game
func newGame(player1 *Player, player2 *Player, points *Score, board [][]string) bool {
	var starter *Player
	var seconder *Player
	var gameOver bool
	occupied := make(map[string]int)

	//toss a coin to determine who's starts
	result := rand.Intn(2)
	if result == 0 {
		starter = player1
		seconder = player2
	} else {
		starter = player2
		seconder = player1
	}

	// for loop for every round of the game
	var i int
	for i = 0; i < 9 || gameOver == false; i++ {
		var nextMove string
		var nextSymbol string
		if i%2 == 0 {
			fmt.Printf("Next player to make a mark is %v \n", starter.name)
			printBoard(board, starter, seconder, points)
			fmt.Printf("Where to put the %v ? (inputr row then column, seperated by space) \n", starter.symbol)
			fmt.Scan(&nextMove)

			// catch invalid input
			valid := checkInput(nextMove, occupied)
			if !valid {
				i--
			} else {
				nextSymbol = starter.symbol
				occupied[nextMove] = i
				rowGiven, _ := strconv.Atoi(nextMove[0:1])
				columnGiven, _ := strconv.Atoi(nextMove[1:2])
				board[rowGiven][columnGiven] = nextSymbol
			}
		} else {
			fmt.Printf("Next player to make a mark is %v \n", seconder.name)
			printBoard(board, starter, seconder, points)
			fmt.Printf("Where to put the %v ? (inputr row then column, seperated by space) \n", seconder.symbol)
			fmt.Scan(&nextMove)

			// catch invalid input
			valid := checkInput(nextMove, occupied)
			if !valid {
				i--
			} else {
				nextSymbol = seconder.symbol
				occupied[nextMove] = i
				rowGiven, _ := strconv.Atoi(nextMove[0:1])
				columnGiven, _ := strconv.Atoi(nextMove[1:2])
				board[rowGiven][columnGiven] = nextSymbol
			}
		}

		res := checkIfWin(board)
		if res {
			if i%2 == 0 {
				fmt.Printf("%v won! click ESC to exit or Space to retry", starter.name)
				points.player1Score++
			} else {
				fmt.Printf("%v won! click ESC to exit or Space to retry", seconder.name)
				points.player2Score++
			}
			gameOver = true
			return gameOver
		}
	}
	if i == 9 {
		points.tide++
	}
	gameOver = true
	return gameOver
}
