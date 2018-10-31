package state

type GameState interface {
	executeState(*GameContext) bool
}

type GameContext struct {
	SecretNumber int
	Retires      int
	Won          bool
	Next         GameState
}
