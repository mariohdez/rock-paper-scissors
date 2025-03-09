package user

import "github.com/mariohdez/rockpaperscissors/internal/model"

type User struct {
	Name   string
	Wins   int
	Weapon model.Weapon
}
