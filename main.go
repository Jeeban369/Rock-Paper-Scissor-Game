//ROCK PAPER SCISSOR GAME

package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// 3 round
// computer vs player
func main() {
	//computer choice logic
	const rounds = 3
	fmt.Println("Rock✊  Paper🫱  Scissor✌️  Game")
	pWinningPoint := 0
	cWinningPoint := 0


	for i := 0; i < rounds; i++ {
		fmt.Printf("🥏Round %v\n",i+1)
		computerChoiceNum := rand.Intn(rounds)

		var computerChoice string

		switch computerChoiceNum {
		case 0:
			computerChoice = "Rock✊"
		case 1:
			computerChoice = "Paper🫱"
		case 2:
			computerChoice = "Scissor✌️"
		}

		//read player choice
		fmt.Print("Enter your choice (Rock, Paper, scissor): ")
		var playerChoice string
		fmt.Scanln(&playerChoice)

		computerChoice = strings.ToLower(computerChoice)
		playerChoice = strings.ToLower(playerChoice)

		//winning logic
		fmt.Printf("Computer choice : %s \n", computerChoice)

		switch {
		case playerChoice == computerChoice:
			fmt.Println("It's a draw.🤝")

		case playerChoice == "rock" && computerChoice == "scissor✌️",
			playerChoice == "paper" && computerChoice == "rock✊",
			playerChoice == "scissor" && computerChoice == "paper🫱":
			fmt.Println("You won the round.✨")
			pWinningPoint++

		default:
			fmt.Println("Computer won the round.🛠️")
			cWinningPoint++

		}
	}
	//game over display
	fmt.Println("Game Over!")
	if pWinningPoint>cWinningPoint{
		fmt.Println("Congrats🎉, You won the match.🥳")
	}else if pWinningPoint == cWinningPoint{
		fmt.Println("It's a draw match.😶‍🌫️")
	}else{
		fmt.Println("Oops🥲, You lose the game. \nBetter Luck Next Time.")
	}
}
