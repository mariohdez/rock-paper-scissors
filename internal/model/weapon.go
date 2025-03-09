package model

import (
	"strconv"
	"strings"
)

type Weapon int32

const (
	UnknownWeapon Weapon = iota
	Rock
	Paper
	Scissors
)

var unknownWeaponStr = "❓"
var weapons = [...]string{unknownWeaponStr, "🪨", "📄", "✂️"}

func (m Weapon) String() string {
	if m < UnknownWeapon || m > Scissors {
		return unknownWeaponStr
	}

	return weapons[m]
}

func ParseWeapon(input string) Weapon {
	input = strings.TrimSpace(strings.ToLower(input))

	if weapon, err := strconv.Atoi(input); err == nil {
		switch weapon {
		case 1:
			return Rock
		case 2:
			return Paper
		case 3:
			return Scissors
		default:
			return UnknownWeapon
		}
	}

	switch input {
	case "rock":
		return Rock
	case "paper":
		return Paper
	case "scissors":
		return Scissors
	}
	return UnknownWeapon
}
