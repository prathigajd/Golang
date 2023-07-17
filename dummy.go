package main

import (
	 "fmt"
	 "math/rand"
)

func main(){

var bankroll float64
fmt.Println("Enter the Intial amount") 
fmt.Scanln(&bankroll)

for bankroll > 0 {
fmt.Println("1. Bet on Red or Black")
fmt.Println("2. Bet on High or Low")
fmt.Println("3. Bet on a Specific Number")
fmt.Println("4. Quit")

var choice int
fmt.Println("Enter the number of your choice")
fmt.Scanln(&choice) 

switch choice {
case 1: 
	 bankroll = betOnColor(bankroll)
case 2: 
     bankroll = betOnRange(bankroll)
case 3 : 
     bankroll = betOnNumber(bankroll)
case 4 :
	 fmt.Println("quit the game")
	 return
default : 
     fmt.Println("Invalid choice")
}

if bankroll == 0 {
	fmt.Println("Ran out of money. Game over")
} else {   
	fmt.Printf("Your current bankroll: $%.2f\n", bankroll)
}
}
}

const (
	red   = "1 3 5 7 9 12 14 16 18 19 21 23 25 27 30 32 34 36"
	black = "2 4 6 8 10 11 13 15 17 20 22 24 26 28 29 31 33 35"
)

func betOnColor(bankroll float64) float64{
     
	var betamt float64
	fmt.Println("Plese enter the amount to bet")
	fmt.Scanln(&betamt)

	if betamt <= 0 || betamt > bankroll {
		fmt.Println("Amount insufficient")
		return bankroll
	}
        var color string
		fmt.Println("choose color")
		fmt.Scanln(&color)

		var winningNumber int = spinWheel()

if color == red || color == black {
			fmt.Println("You won", winningNumber, betamt)
			bankroll += betamt
}else{
			fmt.Println("You lost", winningNumber, betamt)  
			bankroll -= betamt
		}
		return bankroll
	}

	var (
		low   = "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18"
		high = "19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37"
	)

func betOnRange(bankroll float64) float64{
	var betamt float64
	fmt.Println("Plese enter the amount to bet")
	fmt.Scanln(&betamt)

	if betamt<=0 || betamt > bankroll {
		fmt.Println("Amount insufficient")
		return bankroll
	}

	var rangeChoice string
	fmt.Println("choose range (low/high)")
	fmt.Scanln(&rangeChoice)

	var winningNumber int = spinWheel()

		if rangeChoice == low || rangeChoice == high {
			fmt.Println("You won", winningNumber, betamt)
			bankroll += betamt
		}else{
			fmt.Println("You lost", winningNumber, betamt)
			bankroll -= betamt
		}
		return bankroll
}

func betOnNumber(bankroll float64) float64{

	var betamt float64
	fmt.Println("Plese enter the amount to bet")
	fmt.Scanln(&betamt)

	if betamt<=0 || betamt > bankroll {
		fmt.Println("Amount insufficient")
		return bankroll
	}

	var number int
	fmt.Println("Enter the Number")
	fmt.Scanln(&number)

	if number<=0 || number>37 {
		fmt.Println("invalid number")
	}

	var winningNumber int = spinWheel()

		if number == winningNumber {
			fmt.Println("You won", winningNumber, betamt)
			bankroll += betamt*35
		}else{
			fmt.Println("You lost", winningNumber, betamt)
			bankroll -= betamt
		}
		return bankroll
}

func spinWheel() int {
	return rand.Intn(38)
}



