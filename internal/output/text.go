package output

import (
	"fmt"
	"github.com/mariohdez/rockpaperscissors/internal/game"
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
		_, err := fmt.Fprintf(w.writer, "DRAW 🤝⚖️\n\n")
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

func (w *TextResultWriter) WriteMatchOutcome(outcome *game.MatchOutcome) error {
	if outcome.IsDraw {
		_, err := fmt.Fprintf(w.writer, "Great minds think alike 🧠🤝🧠. No winner.\n")
		if err != nil {
			return fmt.Errorf("write outcome to writer: %w", err)
		}
		return nil
	}

	_, err := fmt.Fprintf(w.writer, "%s won! Congratulations! 🎉\n", outcome.WinnerName)
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
