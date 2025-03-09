package game

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mariohdez/rockpaperscissors/internal/model"
	"github.com/mariohdez/rockpaperscissors/internal/user"
)

type RoundWriter interface {
	WriteRound(result *RoundResult)
}

type MatchWriter interface {
	WriteMatch(user1, user2 *user.User)
}

type Match struct {
	user1, user2 *user.User
	roundWriter  RoundWriter
	matchWriter  MatchWriter
	scanner      *bufio.Scanner
	winsNeeded   int
}

func NewMatch(numRounds int, user1, user2 *user.User, scanner *bufio.Scanner, roundWriter RoundWriter, matchWriter MatchWriter) *Match {
	return &Match{
		user1:       user1,
		user2:       user2,
		roundWriter: roundWriter,
		matchWriter: matchWriter,
		scanner:     scanner,
		winsNeeded:  numRounds/2 + 1,
	}
}

func (m *Match) Start() {
	for m.user1.Wins < m.winsNeeded && m.user2.Wins < m.winsNeeded {
		for m.user1.Weapon == model.UnknownWeapon {
			fmt.Printf("%s please enter your move: ", m.user1.Name)
			m.scanner.Scan()
			user1Input := m.scanner.Text()
			m.user1.Weapon = model.ParseWeapon(user1Input)
		}

		for m.user2.Weapon == model.UnknownWeapon {
			fmt.Printf("%s please enter your move: ", m.user2.Name)
			m.scanner.Scan()
			user2Input := m.scanner.Text()
			m.user2.Weapon = model.ParseWeapon(user2Input)
		}

		roundResult, err := NewRoundResult(m.user1, m.user2)
		if err != nil {
			panic(errors.Join(errors.New("build round result"), err))
		}

		if !roundResult.IsDraw {
			if roundResult.Winner == m.user1 {
				m.user1.Wins++
			} else {
				m.user2.Wins++
			}
		}

		m.roundWriter.WriteRound(roundResult)

		m.user1.Weapon = model.UnknownWeapon
		m.user2.Weapon = model.UnknownWeapon
	}

	m.matchWriter.WriteMatch(m.user1, m.user2)
}
