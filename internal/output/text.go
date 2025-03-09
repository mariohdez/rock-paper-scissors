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

func (w *TextResultWriter) WriteRoundOutcome(outcome *game.RoundOutcome) error {
	if outcome.IsDraw {
		_, err := fmt.Fprintln(w.writer, "DRAW 🤝⚖️")
		if err != nil {
			return fmt.Errorf("write outcome to writer: %w", err)
		}
		return nil
	}

	_, err := fmt.Fprintf(
		w.writer,
		"%s beat %s: %s beats %s\n\n",
		outcome.Winner.Name,
		outcome.Loser.Name,
		outcome.Winner.Weapon,
		outcome.Loser.Weapon,
	)
	if err != nil {
		return fmt.Errorf("write outcome to writer: %w", err)
	}

	return nil
}

func (w *TextResultWriter) WriteMatchOutcome(user1, user2 *user.User) error {
	winner := user1
	if user2.Wins > user1.Wins {
		winner = user2
	}

	_, err := fmt.Fprintf(w.writer, "%s won! Congratulations! 🎉\n", winner.Name)
	if err != nil {
		return fmt.Errorf("write outcome to writer: %w", err)
	}

	return nil
}

func (w *TextResultWriter) WriteMatchError(err error) error {
	_, writeErr := fmt.Fprintln(w.writer, "The match encountered an error:", err)
	if writeErr != nil {
		return fmt.Errorf("write outcome to writer: %w", err)
	}

	return nil
}
