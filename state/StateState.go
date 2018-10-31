package state

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type StartState struct{}

func (s *StartState) executeState(c *GameContext) bool {
	c.Next = &AskState{}

	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(10)

	fmt.Println("Introduce a number a number of retries to set the difficulty")
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retires)

	return true
}
