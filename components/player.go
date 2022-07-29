package main

import (
	"fmt"
	"strings"
)

type iPlayer interface {
	setSymbol() string
	getSymbol() string
	getNextMove() (int, int, error)
	getScore() int
	getID() int
}

type Player struct {
	name   string
	symbol string
	score  int
}

func newPlayer(symbol string, name string) *Player {
	p := Player{
		name:   name,
		symbol: symbol,
		score:  0,
	}
	return &p
}

func setPlayers() (*Player, *Player) {
	var player1Sign string
	var player1Name string
	var player1SignNumber string
	var player2Sign string
	var player2Name string

	fmt.Println("Choose first player's name")
	fmt.Scan(&player1Name)
	fmt.Printf("%v choose your sign: 'X - 1' or 'O - 2'", player1Name)
	fmt.Scan(&player1SignNumber)
	fmt.Println("Choose second player's name")
	fmt.Scan(&player2Name)

	if player1SignNumber == "1" {
		player1Sign = "X"
	} else if player1SignNumber == "2" {
		player1Sign = "Y"
	} else {
		fmt.Println("invalid input")
	}

	player1 := newPlayer(player1Sign, player1Name)

	if strings.ToLower(player1Sign) == "O" {
		player2Sign = "X"
	} else {
		player2Sign = "O"
	}

	player2 := newPlayer(player2Sign, player2Name)

	return player1, player2

}
