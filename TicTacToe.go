package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var board [3][3]string = [3][3]string{[3]string{" ", " ", " "}, [3]string{" ", " ", " "}, [3]string{" ", " ", " "}}
var keymap map[string]string = map[string]string{
	" ": " ",
	"0": "0",
	"x": "x",
}
var player string = "x"
var row int
var column int64

func show_board() {
	fmt.Printf("\n|%v|%v|%v|\t\t|1|2|3|", board[0][0], board[0][1], board[0][2])
	fmt.Printf("\n|%v|%v|%v|\t\t|4|5|6|", board[1][0], board[1][1], board[1][2])
	fmt.Printf("\n|%v|%v|%v|\t\t|7|8|9|", board[2][0], board[2][1], board[2][2])
}

func _check_wins() bool {
	p := keymap[player]
	for _, i := range board {
		if i[0] == p && i[1] == p && i[2] == p {
			return true
		}
	}

	for j := 0; j < len(board); j++ {
		if board[0][j] == p && board[1][j] == p && board[2][j] == p {
			return true
		}
	}

	if board[0][0] == p && board[1][1] == p && board[2][2] == p {
		return true
	}

	if board[0][2] == p && board[1][1] == p && board[2][0] == p {
		return true
	}

	return false
}

func check_wins() {
	if _check_wins() {
		show_board()
		fmt.Printf("\n\n%v WON!\n", player)
		os.Exit(0)
	}
}

func _check_ends() bool {
	for _, i := range board {
		for _, j := range i {
			if j == " " {
				return false
			}
		}
	}

	return true
}

func check_ends() {
	if _check_ends() {
		show_board()
		fmt.Println("\n\nGame has been ended and no one won!")
		os.Exit(0)
	}
}

func press() {
	show_board()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\n\n%v - what position do you choose: ", player)
	scanner.Scan()
	inputstring := scanner.Text()
	input, _ := strconv.ParseInt(inputstring, 10, 64)
	input -= 1

	switch input {
	case 0, 1, 2:
		row = 0
		column = input

	case 3, 4, 5:
		row = 1
		column = input - 3

	case 6, 7, 8:
		row = 2
		column = input - 6
	}

	if board[row][column] != " " {
		fmt.Println("\nInvalid choice")
		press()
	} else {
		board[row][column] = player
		check_wins()
		check_ends()

		if player == "x" {
			player = "0"
		} else {
			player = "x"
		}
	}
}

func main() {
	for {
		press()
	}
}
