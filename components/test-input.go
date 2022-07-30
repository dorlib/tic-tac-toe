package main

import "fmt"

func checkInput(input string, occupied map[string]int) bool {

	if len(input) != 2 {
		fmt.Printf("***invalid input : length of input must be 2 \n")
		return false
	} else if (input[0:1] != "1" && input[0:1] != "2" && input[0:1] != "0") || (input[1:] != "1" && input[1:] != "2" && input[1:] != "0") {
		fmt.Printf("***invalid input : numbers not in range \n")
		return false
	} else if occupied[input] != 0.0 {
		fmt.Printf("***invalid input : given place is occupied \n")
		return false
	}
	return true
}
