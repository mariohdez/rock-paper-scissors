package view

import (
	"fmt"
	"github.com/mariohdez/rockpaperscissors/internal/game"
	"github.com/mariohdez/rockpaperscissors/internal/user"
	"io"
)

type TextResultWriter struct {
	writer io.Writer
}

func New(w io.Writer) *TextResultWriter {
	return &TextResultWriter{
		writer: w,
	}
}

func (w *TextResultWriter) WriteRound(result *game.RoundResult) {
	if result.IsDraw {
		_, _ = fmt.Fprintln(w.writer, "DRAW 🤝⚖️")
		return
	}

	_, _ = fmt.Fprintf(
		w.writer,
		"%s beat %s: %s beats %s\n\n",
		result.Winner.Name,
		result.Loser.Name,
		result.Winner.Weapon,
		result.Loser.Weapon,
	)
}

func (w *TextResultWriter) WriteMatch(user1, user2 *user.User) {
	winner := user1
	if user2.Wins > user1.Wins {
		winner = user2
	}

	_, _ = fmt.Fprintf(w.writer, "%s won! Congratulations! 🎉\n", winner.Name)
}
