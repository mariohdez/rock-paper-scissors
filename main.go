package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mariohdez/rockpaperscissors/internal/display"
	"github.com/mariohdez/rockpaperscissors/internal/game"
	"github.com/mariohdez/rockpaperscissors/internal/model"
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

	output := display.New(os.Stdout)
	reader := bufio.NewScanner(os.Stdin)
	for user1.Wins < numWins && user2.Wins < numWins {
		for user1.Weapon == model.UnknownWeapon {
			fmt.Printf("%s please enter your move: ", user1.Name)
			reader.Scan()
			user1Input := reader.Text()
			user1.Weapon = model.ToWeapon(user1Input)
		}

		for user2.Weapon == model.UnknownWeapon {
			fmt.Printf("%s please enter your move: ", user2.Name)
			reader.Scan()
			user2Input := reader.Text()
			user2.Weapon = model.ToWeapon(user2Input)
		}

		roundResult, err := game.NewRoundResult(user1, user2)
		if err != nil {
			panic(errors.Join(errors.New("build round result"), err))
		}

		if !roundResult.IsDraw {
			if roundResult.Winner == user1 {
				user1.Wins++
			} else {
				user2.Wins++
			}
		}

		output.RoundResult(roundResult)

		// reset the moves...
		user1.Weapon = model.UnknownWeapon
		user2.Weapon = model.UnknownWeapon
	}

	output.GameResult(user1, user2)
}
