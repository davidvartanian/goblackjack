package main

import (
	"fmt"
	"github.com/davidvartanian/goblackjack/pkg"
	"github.com/davidvartanian/godeck"
	"os"
)

func main() {
	var gs pkg.GameState
	gs = Shuffle(gs)
	gs = Deal(gs)

	var input string
	for gs.State == pkg.StatePlayerTurn {
		fmt.Println("Player:", gs.Player)
		fmt.Println("Dealer:", gs.Dealer.DealerString())
		fmt.Println("What do you wanna do? (h)it, (s)tand, (e)xit")
		_, _ = fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			fmt.Println("-- HIT --")
			gs = Hit(gs)
		case "s":
			fmt.Println("stand too")
			gs = Stand(gs)
		case "e":
			fmt.Println("-- BYE --")
			os.Exit(0)
		default:
			fmt.Println("Invalid option:", input)
		}
	}

	// dealer logic
	for gs.State == pkg.StateDealerTurn {
		if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}
	EndGame(gs)
}

func draw(cards []godeck.Card) (godeck.Card, []godeck.Card) {
	return cards[0], cards[1:]
}

func Shuffle(gs pkg.GameState) pkg.GameState {
	ret := pkg.Clone(gs)
	ret.Deck = godeck.New(godeck.Deck(3), godeck.Shuffle)
	return ret
}

func Deal(gs pkg.GameState) pkg.GameState {
	ret := pkg.Clone(gs)
	ret.Player = make(pkg.Hand, 0, 5)
	ret.Dealer = make(pkg.Hand, 0, 5)
	var card godeck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = pkg.StatePlayerTurn
	return ret
}

func Hit(gs pkg.GameState) pkg.GameState {
	ret := pkg.Clone(gs)
	hand := ret.CurrentPlayer()
	var card godeck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

func Stand(gs pkg.GameState) pkg.GameState {
	ret := pkg.Clone(gs)
	ret.State++
	return ret
}

func EndGame(gs pkg.GameState) pkg.GameState {
	ret := pkg.Clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("-- FINAL HAND --")
	fmt.Println("Player:", ret.Player, "Score:", pScore)
	fmt.Println("Dealer:", ret.Dealer, "Score:", dScore)

	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose")
	default:
		fmt.Println("Draw")
	}
	ret.Player = nil
	ret.Dealer = nil
	return ret
}
