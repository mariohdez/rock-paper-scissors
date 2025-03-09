package game

import (
	"github.com/mariohdez/rockpaperscissors/internal/model"
	"github.com/mariohdez/rockpaperscissors/internal/user"
)

type RoundWriter interface {
	WriteRoundOutcome(result *RoundOutcome) error
}

type MatchWriter interface {
	WriteMatchOutcome(user1, user2 *user.User) error
}

type WeaponReader interface {
	ReadWeapon(user *user.User) error
}

type Match struct {
	user1, user2 *user.User
	roundWriter  RoundWriter
	matchWriter  MatchWriter
	weaponReader WeaponReader
	winsNeeded   int
}

func NewMatch(numRounds int, user1, user2 *user.User, weaponReader WeaponReader, roundWriter RoundWriter, matchWriter MatchWriter) *Match {
	return &Match{
		user1:        user1,
		user2:        user2,
		weaponReader: weaponReader,
		roundWriter:  roundWriter,
		matchWriter:  matchWriter,
		winsNeeded:   numRounds/2 + 1,
	}
}

func (m *Match) Start() error {
	for m.user1.Wins < m.winsNeeded && m.user2.Wins < m.winsNeeded {
		err := m.weaponReader.ReadWeapon(m.user1)
		if err != nil { // return error here?
			return err
		}

		err = m.weaponReader.ReadWeapon(m.user2)
		if err != nil {
			return err
		}

		outcome, err := NewRoundOutcome(m.user1, m.user2)
		if err != nil {
			return err
		}

		m.updateWinner(outcome)
		err = m.roundWriter.WriteRoundOutcome(outcome)
		if err != nil {
			return err
		}

		m.resetRound()
	}

	err := m.matchWriter.WriteMatchOutcome(m.user1, m.user2)
	if err != nil {
		return err
	}

	return nil
}

func (m *Match) updateWinner(outcome *RoundOutcome) {
	if outcome.IsDraw {
		return
	}

	if outcome.Winner == m.user1 {
		m.user1.Wins++
	} else {
		m.user2.Wins++
	}
}

func (m *Match) resetRound() {
	m.user1.Weapon = model.UnknownWeapon
	m.user2.Weapon = model.UnknownWeapon
}
