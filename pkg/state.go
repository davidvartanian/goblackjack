package pkg

import "github.com/davidvartanian/godeck"

type State uint8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

// GameState holds the game state data
type GameState struct {
	Deck   []godeck.Card
	State  State
	Player Hand
	Dealer Hand
}

// CurrentPlayer returns the current player according to the game state
func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("no current player's turn")
	}
}
