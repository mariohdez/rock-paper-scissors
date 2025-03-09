package main

import (
	"fmt"
)

type user struct {
	wins        int
	currentMove int
}

func main() {
	numRounds := 3
	numWins := numRounds/2 + 1

	fmt.Println(numRounds)
	fmt.Println(numWins)

	user1 := &user{}
	user2 := &user{}

	for user1.wins < numWins && user2.wins < numWins {
		// user 1 will get input from stdin.
		// then user2 will get input based on the outcome of random generator.

		// I'll compare to see who won, then round, it's possible there will be a tie.

		// I want to "draw" the round outcome with emojis to make it look cool
	}

	// here I'll print out who one
	// maybe i can have a "wanna keep going?" button.
}
