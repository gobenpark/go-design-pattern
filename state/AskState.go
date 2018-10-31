package state

import (
	"fmt"
	"os"
)

type AskState struct{}

func (a *AskState) executeState(c *GameContext) bool {
	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left \n", c.Retires)

	var n int
	fmt.Fscanf(os.Stdin, "%d", &n)
	c.Retires = c.Retires - 1

	if n == c.SecretNumber {
		c.Won = true
		c.Next = &FinishState{}
	}

	if c.Retires == 0 {
		c.Next = &FinishState{}
	}

	return true
}
