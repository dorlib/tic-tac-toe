package main

import "fmt"

func checkInput(input string, occupied map[string]int) bool {
	if len(input) != 2 {
		fmt.Println("invalid input : length of input must be 2")
		return false
	} else if (input[0:1] != "1" && input[0:1] != "2" && input[0:1] != "0") || (input[1:] != "1" && input[1:] != "2" && input[1:] != "0") {
		fmt.Println("invalid input : numbers not in range")
		return false
	} else if occupied[input] != 0.0 {
		fmt.Println("invalid input : given place is occupied")
		return false
	}
	return true
}
