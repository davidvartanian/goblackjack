package pkg

import (
	"fmt"
	"github.com/davidvartanian/godeck"
	"strings"
)

// Hand is the set of cards given to a player
type Hand []godeck.Card

// String representation of the hand
func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

// DealerString is the string representation of the dealer's hand
func (h Hand) DealerString() string {
	return fmt.Sprintf("%s, ***HIDDEN***", h[0].String())
}

// MinScore returns the minimum hand's score
func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

// Score returns the actual hand's score
func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == godeck.Ace {
			return minScore + 10
		}
	}
	return minScore
}
