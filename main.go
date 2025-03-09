package main

import (
	"bufio"
	"github.com/mariohdez/rockpaperscissors/internal/game"
	"github.com/mariohdez/rockpaperscissors/internal/input"
	"github.com/mariohdez/rockpaperscissors/internal/output"
	"github.com/mariohdez/rockpaperscissors/internal/user"
	"os"
)

func main() {
	user1 := &user.User{
		Name: "Rob",
	}
	user2 := &user.User{
		Name: "Bob",
	}

	writer := os.Stdout
	textOutcomeWriter := output.NewTextOutcomeWriter(writer)
	match := game.NewMatch(
		3,
		user1,
		user2,
		input.NewTextReader(bufio.NewScanner(os.Stdin), writer),
		textOutcomeWriter,
		textOutcomeWriter,
	)

	err := match.Start()
	if err != nil {
		textOutcomeWriter.WriteMatchError(err)
	}
}
