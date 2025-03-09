package user

import "github.com/mariohdez/rockpaperscissors/internal/model"

type Player struct {
	Name   string
	Wins   int
	Weapon model.Weapon
}
