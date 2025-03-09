package display

import (
	"fmt"
	"github.com/mariohdez/rockpaperscissors/internal/game"
	"github.com/mariohdez/rockpaperscissors/internal/user"
	"io"
)

type Display struct {
	writer io.Writer
}

func New(w io.Writer) *Display {
	return &Display{
		writer: w,
	}
}

func (d *Display) RoundResult(result *game.RoundResult) {
	if result.IsDraw {
		_, _ = fmt.Fprintf(d.writer, "DRAW 🤝⚖️\n")
		return
	}

	_, _ = fmt.Fprintf(
		d.writer,
		"%s beat %s: %s beats %s\n\n",
		result.Winner.Name,
		result.Loser.Name,
		result.Winner.Weapon,
		result.Loser.Weapon,
	)
}

func (d *Display) GameResult(user1, user2 *user.User) {
	winner := user1
	if user2.Wins > user1.Wins {
		winner = user2
	}

	_, _ = fmt.Fprintf(d.writer, "%s won! Congratulations! 🎉\n", winner.Name)
}
