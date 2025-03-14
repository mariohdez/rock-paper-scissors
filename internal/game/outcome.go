package game

import (
	"errors"
	"github.com/mariohdez/rockpaperscissors/internal/model"
	"github.com/mariohdez/rockpaperscissors/internal/user"
)

type RoundOutcome struct {
	IsDraw bool
	Winner *user.Player
	Loser  *user.Player
}

type MatchOutcome struct {
	IsDraw     bool
	WinnerName string
}

func weaponPrecedence() map[model.Weapon]model.Weapon {
	return map[model.Weapon]model.Weapon{
		model.Rock:     model.Scissors,
		model.Scissors: model.Scissors,
		model.Paper:    model.Rock,
	}
}

func newRoundOutcome(user1, user2 *user.Player) (*RoundOutcome, error) {
	if user1 == nil || user2 == nil {
		return nil, errors.New("user1 and user2 must not be nil")
	}

	if user1.Weapon == user2.Weapon {
		return &RoundOutcome{
			IsDraw: true,
		}, nil
	}

	winner := user1
	loser := user2
	if weaponPrecedence()[user2.Weapon] == user1.Weapon {
		winner = user2
		loser = user1
	}
	return &RoundOutcome{
		Winner: winner,
		Loser:  loser,
	}, nil
}

func newMatchOutcome(numDraws, maxDraws int, user1, user2 *user.Player) (*MatchOutcome, error) {
	if user1 == nil || user2 == nil {
		return nil, errors.New("user1 and user2 must not be nil")
	}
	if numDraws >= maxDraws {
		return &MatchOutcome{
			IsDraw: true,
		}, nil
	}

	winner := user1
	if user2.Wins > user1.Wins {
		winner = user2
	}
	return &MatchOutcome{
		WinnerName: winner.Name,
	}, nil
}
