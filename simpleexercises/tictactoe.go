package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkforWin(board [][]int) (bool, bool) {
	counter := 0
	gameOver := false
	for i, row := range board {
		rowplayer1, rowplayer2 := 0, 0
		colplayer1, colplayer2 := 0, 0
		for j, cell := range row {

			//checking rows
			if cell == 1 {
				rowplayer1 = rowplayer1 + 1
				counter = counter + 1
			} else if cell == 2 {
				rowplayer2 = rowplayer2 + 1
				counter = counter + 1
			}

			//check columns
			if board[j][i] == 1 {
				colplayer1 = colplayer1 + 1
			} else if board[j][1] == 2 {
				colplayer2 = colplayer2 + 1
			}
		}

		if counter == 9 {
			gameOver = true
		}

		if rowplayer1 == 3 || rowplayer2 == 3 || colplayer1 == 3 || colplayer2 == 3 {

			return true, gameOver
		}
	}

	//check diagonals
	dg1p1, dg1p2 := 0, 0
	dg2p1, dg2p2 := 0, 0
	size := len(board[0])
	for d := 0; d < size; d++ {
		if board[d][d] == 1 {
			dg1p1 = dg1p1 + 1
		} else if board[d][d] == 2 {
			dg1p2 = dg1p2 + 1
		}

		if board[d][size-1-d] == 1 {
			dg2p1 = dg2p1 + 1
		} else if board[d][size-1-d] == 2 {
			dg2p2 = dg2p2 + 1
		}
	}

	if dg1p1 == 3 || dg1p2 == 3 || dg2p1 == 3 || dg2p2 == 3 {
		return true, gameOver
	}

	return false, gameOver
}

func takeInputfromUser(player int) (int, int) {
	fmt.Printf("Player %d's turn ", player)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if scanner.Err() != nil {
		fmt.Println("Invalid Input, Game Over")
	}
	choice := scanner.Text()
	slc := strings.Fields(choice)

	x, _ := strconv.Atoi(slc[0])
	y, _ := strconv.Atoi(slc[1])

	return x, y
}

func initializeBoard(size int) [][]int {
	board := make([][]int, size)

	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}

	return board
}

func updateBoard(board [][]int, x, y, player int) (bool, bool) {
	board[x][y] = player
	win, over := checkforWin(board)
	return win, over
}

func boardslotoccupied(board [][]int, x, y int) bool {
	if board[x][y] != 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Tic Tac Toe Game")
	board := initializeBoard(3)
	fmt.Println("Please input the position number of your move as coordinates X Y")
	player := 1
	for {
		x, y := takeInputfromUser(player)
		if x > 2 || y > 2 || x < 0 || y < 0 {
			fmt.Println("Value out of range")
			continue
		}
		if boardslotoccupied(board, x, y) {
			fmt.Println("Can't override")
			continue
		}
		won, over := updateBoard(board, x, y, player)
		for _, row := range board {
			fmt.Println(row)
		}

		if won == true {
			fmt.Printf("Game over, %d player won the game", player)
			break
		} else if over == true {
			fmt.Println("Game Over, Tie!!")
			break
		} else {
			if player == 1 {
				player = 2
			} else {
				player = 1
			}
		}
	}
}
