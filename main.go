package main

import (
	"bufio"
	"github.com/mariohdez/rockpaperscissors/internal/game"
	"github.com/mariohdez/rockpaperscissors/internal/user"
	"github.com/mariohdez/rockpaperscissors/internal/view"
	"os"
)

func main() {
	user1 := &user.User{
		Name: "Rob",
	}
	user2 := &user.User{
		Name: "Bob",
	}

	writer := view.New(os.Stdout)
	match := game.NewMatch(
		3,
		user1,
		user2,
		bufio.NewScanner(os.Stdin),
		writer,
		writer,
	)
	match.Start()
}
