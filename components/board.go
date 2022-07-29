package main

import "fmt"

func initBoard() [][]string {
	return [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
}

func printBoard(board [][]string, player1 *Player, player2 *Player, tide int) {
	fmt.Printf("  %v  |  %v  |  %v  ", board[0][0], board[0][1], board[0][2])
	fmt.Printf("     |  %v's score: %v  |\n", player1.name, player1.score)
	fmt.Printf("________________")
	fmt.Printf("      |  %v's score: %v  |\n", player2.name, player2.score)
	fmt.Printf("  %v  |  %v  |  %v  ", board[1][0], board[1][1], board[1][2])
	fmt.Printf("     |  tide: %v          |\n", tide)
	fmt.Printf("________________")
	fmt.Printf("      |  now play's: %v  |\n", player1.name)
	fmt.Printf("  %v  |  %v  |  %v  ", board[2][0], board[2][1], board[2][2])
}

func checkIfWin(board [][]string) bool {
	switch {
	case board[0][0] == board[0][1] && board[0][1] == board[0][2]:
		return true
	case board[1][0] == board[1][1] && board[1][1] == board[1][2]:
		return true
	case board[2][0] == board[2][1] && board[2][1] == board[2][2]:
		return true
	case board[0][0] == board[1][0] && board[1][0] == board[2][0]:
		return true
	case board[0][1] == board[1][1] && board[1][1] == board[2][1]:
		return true
	case board[0][2] == board[1][2] && board[1][2] == board[2][2]:
		return true
	case board[0][0] == board[1][1] && board[1][1] == board[2][2]:
		return true
	case board[0][2] == board[1][1] && board[1][1] == board[2][0]:
		return true
	default:
		return false
	}
}
