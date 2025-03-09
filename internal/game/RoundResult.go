package game

import (
	"errors"
	"github.com/mariohdez/rockpaperscissors/internal/model"
	"github.com/mariohdez/rockpaperscissors/internal/user"
)

type RoundResult struct {
	IsDraw bool
	Winner *user.User
	Loser  *user.User
}

func weaponPrecedence() map[model.Weapon]model.Weapon {
	return map[model.Weapon]model.Weapon{
		model.Rock:     model.Scissors,
		model.Scissors: model.Scissors,
		model.Paper:    model.Rock,
	}
}

func NewRoundResult(user1, user2 *user.User) (*RoundResult, error) {
	if user1 == nil || user2 == nil {
		return nil, errors.New("user1 and user2 must not be nil")
	}

	if user1.Weapon == user2.Weapon {
		return &RoundResult{
			IsDraw: true,
		}, nil
	}

	winner := user1
	loser := user2
	if weaponPrecedence()[user2.Weapon] == user1.Weapon {
		winner = user2
		loser = user1
	}
	return &RoundResult{
		Winner: winner,
		Loser:  loser,
	}, nil
}
