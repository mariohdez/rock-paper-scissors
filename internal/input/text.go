package input

import (
	"bufio"
	"fmt"
	"github.com/mariohdez/rockpaperscissors/internal/model"
	"github.com/mariohdez/rockpaperscissors/internal/user"
	"io"
)

type TextInputReader struct {
	scanner *bufio.Scanner
	writer  io.Writer
}

func NewTextReader(scanner *bufio.Scanner, writer io.Writer) *TextInputReader {
	return &TextInputReader{
		scanner: scanner,
		writer:  writer,
	}
}

func (t *TextInputReader) ReadWeapon(user *user.Player) error {
	const maxAttempts = 3
	numOfInvalidAttempts := 0
	for user.Weapon == model.UnknownWeapon && numOfInvalidAttempts < maxAttempts {
		_, err := fmt.Fprintf(t.writer, "%s, please enter your weapon: ", user.Name)
		if err != nil {
			return fmt.Errorf("write prompt to read weapon: %w", err)
		}

		if !t.scanner.Scan() {
			if err := t.scanner.Err(); err != nil {
				return fmt.Errorf("read from scanner: %w", t.scanner.Err())
			}
			return fmt.Errorf("scanner reached EOF")
		}
		userInput := t.scanner.Text()
		user.Weapon = model.ParseWeapon(userInput)
		if user.Weapon == model.UnknownWeapon {
			numOfInvalidAttempts++
		}
	}

	if numOfInvalidAttempts >= maxAttempts {
		return fmt.Errorf("too many invalid inputs")
	}
	return nil
}
