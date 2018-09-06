package main

import (
	"fmt"
	"math/rand"
)

const (
	ROCK int = iota
	PAPER
	SCISSORS
)

type Choice struct {
	Who   int //0 you 1 your opponent
	Guess int
}

//Win returns true if you win.
func Win(you, he int) bool {
	if you == ROCK && he == SCISSORS {
		return true
	}
	if you == PAPER && he == ROCK {
		return true
	}
	if you == SCISSORS && he == PAPER {
		return true
	}
	return false
}

func Opponent(guess chan Choice, please chan struct{}) {
	for i := 0; i < 3; i++ {
		<-please
		choice := rand.Intn(3)
		who := 1
		guess <- Choice{who, choice}
		please <- struct{}{}
	}
}

var Cheat func(guess chan Choice) chan Choice

//cheat function is reading the choices and then changing the players choice after peeping the choice of the opponent
//to ensure that player always wins
func cheat(guess chan Choice) chan Choice {
	new_guess := make(chan Choice)
	go func() {
	for i := 0; i < 3; i++ {
		g1 := <-guess
		g2 := <-guess

		if g1.Who == 0 {
			switch g2.Guess {
			case ROCK:
				g1.Guess = PAPER
			case PAPER:
				g1.Guess = SCISSORS
			case SCISSORS:
				g1.Guess = ROCK
			}
			new_guess <- g1
			new_guess <- g2
		} else {
			switch g1.Guess {
			case ROCK:
				g2.Guess = PAPER
			case PAPER:
				g2.Guess = SCISSORS
			case SCISSORS:
				g2.Guess = ROCK
			}
			new_guess <- g2
			new_guess <- g1
		}
	}
	}()
	return new_guess
	}

func Me(guess chan Choice, please chan struct{}) {
	for i := 0; i < 3; i++ {
		<-please
		choice := rand.Intn(3)
		who := 0
		guess <- Choice{who, choice}
		please <- struct{}{}
	}
}

func Game() []bool {
	guess := make(chan Choice)
	//please sync 2 goroutines.
	please := make(chan struct{})
	go func() { please <- struct{}{} }()
	go Opponent(guess, please)
	go Me(guess, please)
	guess = Cheat(guess)
	var wins []bool

	for i := 0; i < 3; i++ {
		g1 := <-guess
		g2 := <-guess
		win := false
		if g1.Who == 0 {
			win = Win(g1.Guess, g2.Guess)
		} else {
			win = Win(g2.Guess, g1.Guess)
		}
		wins = append(wins, win)
	}

	return wins
}

func init() {
	Cheat = cheat
}

func main() {
	win := Game()
	for _, w := range win {
		if !w {
			fmt.Println("Failed!")
		} else {
			fmt.Println("Won!")
		}
	}
	
}