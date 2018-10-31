package state

import "fmt"

type FinishState struct{}

func (f *FinishState) executeState(c *GameContext) bool {
	if c.Won {
		fmt.Println("congrates, you won")
	} else {
		fmt.Println("You lose")
	}
	return false
}
