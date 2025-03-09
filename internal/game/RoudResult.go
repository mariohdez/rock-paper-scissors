package game

import "github.com/mariohdez/rockpaperscissors/internal/user"

type RoundResult struct {
	IsDraw bool
	Winner *user.User
	Loser  *user.User
}

var movePrecedence = map[user.Move]user.Move{
	user.Rock:     user.Scissors,
	user.Scissors: user.Paper,
	user.Paper:    user.Rock,
}

func NewRoundResult(user1, user2 *user.User) *RoundResult {
	if user1.Move == user2.Move {
		return &RoundResult{
			IsDraw: true,
		}
	}

	if movePrecedence[user1.Move] == user2.Move {
		return &RoundResult{
			Winner: user1,
			Loser:  user2,
		}
	}
	return &RoundResult{
		Winner: user2,
		Loser:  user1,
	}
}
