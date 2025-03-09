package output

import (
	"fmt"
	"github.com/mariohdez/rockpaperscissors/internal/game"
	"github.com/mariohdez/rockpaperscissors/internal/user"
	"io"
)

type TextResultWriter struct {
	writer io.Writer
}

func NewTextOutcomeWriter(w io.Writer) *TextResultWriter {
	return &TextResultWriter{
		writer: w,
	}
}

func (w *TextResultWriter) WriteRoundOutcome(outcome *game.RoundOutcome) {
	if outcome.IsDraw {
		_, _ = fmt.Fprintln(w.writer, "DRAW 🤝⚖️")
		return
	}

	_, _ = fmt.Fprintf(
		w.writer,
		"%s beat %s: %s beats %s\n\n",
		outcome.Winner.Name,
		outcome.Loser.Name,
		outcome.Winner.Weapon,
		outcome.Loser.Weapon,
	)
}

func (w *TextResultWriter) WriteMatchOutcome(user1, user2 *user.User) {
	winner := user1
	if user2.Wins > user1.Wins {
		winner = user2
	}

	_, _ = fmt.Fprintf(w.writer, "%s won! Congratulations! 🎉\n", winner.Name)
}

func (w *TextResultWriter) WriteMatchError(err error) {
	_, _ = fmt.Fprintln(w.writer, "The match encountered an error:", err)
}
