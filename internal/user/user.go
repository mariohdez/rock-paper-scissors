package user

import (
	"strconv"
	"strings"
)

type Move int32

const (
	UnknownMove Move = iota
	Rock
	Paper
	Scissors
)

var moves = [...]string{"❓", "🪨", "📄", "✂️"}

func (m Move) String() string {
	if m < UnknownMove || m > Scissors {
		return moves[0]
	}

	return moves[m]
}

func ToMove(input string) Move {
	input = strings.TrimSpace(input)
	if len(input) > 1 {
		return UnknownMove
	}

	move, err := strconv.Atoi(input)
	if err != nil {
		return UnknownMove
	}
	if move < 1 || move > 3 {
		return UnknownMove
	}

	return Move(move)
}

type User struct {
	Name string
	Wins int
	Move Move
}
