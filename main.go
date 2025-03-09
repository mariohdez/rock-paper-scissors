package main

import (
	"bufio"
	"fmt"
	"github.com/mariohdez/rockpaperscissors/internal/display"
	"github.com/mariohdez/rockpaperscissors/internal/game"
	"github.com/mariohdez/rockpaperscissors/internal/user"
	"os"
)

func main() {
	numRounds := 3
	numWins := numRounds/2 + 1

	user1 := &user.User{
		Name: "Rob",
	}
	user2 := &user.User{
		Name: "Bob",
	}

	display := display.New(os.Stdout)
	reader := bufio.NewScanner(os.Stdin)
	for user1.Wins < numWins && user2.Wins < numWins {
		for user1.Move == user.UnknownMove {
			fmt.Printf("%s please enter your move: ", user1.Name)
			reader.Scan()
			user1Input := reader.Text()
			user1.Move = user.ToMove(user1Input)
		}

		for user2.Move == user.UnknownMove {
			fmt.Printf("%s please enter your move: ", user2.Name)
			reader.Scan()
			user2Input := reader.Text()
			user2.Move = user.ToMove(user2Input)
		}

		roundResult := game.NewRoundResult(user1, user2)
		if !roundResult.IsDraw {
			if roundResult.Winner == user1 {
				user1.Wins++
			} else {
				user2.Wins++
			}
		}

		// display the game
		display.Result(roundResult)

		// reset the moves...
		user1.Move = user.UnknownMove
		user2.Move = user.UnknownMove
	}

	display.GameResult(user1, user2)
}
