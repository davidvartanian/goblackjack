package pkg

import "github.com/davidvartanian/godeck"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Clone a game state
func Clone(state GameState) GameState {
	ret := GameState{
		Deck:   make([]godeck.Card, len(state.Deck)),
		State:  state.State,
		Player: make(Hand, len(state.Player)),
		Dealer: make(Hand, len(state.Dealer)),
	}
	copy(ret.Deck, state.Deck)
	copy(ret.Player, state.Player)
	copy(ret.Dealer, state.Dealer)
	return ret
}
