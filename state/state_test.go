package state

import "testing"

func TestState(t *testing.T) {
	start := StartState{}
	game := GameContext{
		Next: &start,
	}
	for game.Next.executeState(&game) {
	}

}
