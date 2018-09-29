package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func checkforWin(board [][]int) bool {

	for i, row := range board {
		rowplayer1, rowplayer2 := 0, 0
		colplayer1, colplayer2 := 0, 0
		for j, cell := range row {

			//checking rows
			if cell == 1 {
				rowplayer1 = rowplayer1 + 1
			} else if cell == 2 {
				rowplayer2 = rowplayer2 + 1
			}

			//check columns
			if board[j][i] == 1 {
				colplayer1 = colplayer1 + 1
			} else if board[j][1] == 2 {
				colplayer2 = colplayer2 + 1
			}
		}
		if rowplayer1 == 3 || rowplayer2 == 3 || colplayer1 == 3 || colplayer2 == 3 {
			return true
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
		return true
	}

	return false
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

func initializeBoard(size int) [][]int{
	board := make([][]int, size)
	
	for i:=0; i<size; i++ {
		board[i] = make([]int, size)
	}
	
	return board
}

func updateBoard(board [][]int, x, y, player int) bool {
	board[x][y] = player
	return checkforWin(board)
}

func main() {
	fmt.Println("Tic Tac Toe Game")
	board := initializeBoard(3)
	fmt.Println("Please input the position number of your move as coordinates X Y")
	player := 1
	for {
	x, y := takeInputfromUser(player)
	won := updateBoard(board, x, y, player)
	fmt.Println(board)

	if won == true {
		fmt.Printf("Game over, %d player won the game", player)
		break;
	} else {
		if player == 1 {
			player = 2
		} else {
			player = 1
		}
	}
	}
}