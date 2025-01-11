package play_hand

import (
	"fmt"

	initPac "github.com/thomasZimmerman0/blackjack_simultor/initialization"
)

func DealCard(game *initPac.Game) initPac.Card{
	card, shoe := game.Shoe.Cards[0], game.Shoe.Cards[1:len(game.Shoe.Cards)]
	game.Shoe.Cards = shoe
	return card
}

func CleanTable(game *initPac.Game, player *initPac.Player){
	game.DealersHand = []initPac.Card{}
	game.Upcard = initPac.Card{Name: "None", Value: 0, Suit: "None"}
	player.Hand = []initPac.Card{}
	player.Bet = 0
}

func StartHand(game *initPac.Game, player *initPac.Player){
	player.Hand = append(player.Hand, DealCard(game))
	game.DealersHand = append(game.DealersHand, DealCard(game))
	player.Hand = append(player.Hand, DealCard(game))
	upcard := DealCard(game)
	game.DealersHand = append(game.DealersHand, upcard)
	game.Upcard = upcard
}

// func CalculateBasicStrategy(game *initPac.Game, player *initPac.Player) string {

// }

// Determines the value of the hand accounting for all the aces ( reducing as many to 1 as necessary )
func DetermineHandValueInFull(hand []initPac.Card) initPac.HandWithValue{
	value := DeterminehandValueSimple(hand)
	if value > 21 {
		acesAtEleven := []*initPac.Card{}
		for i := range hand {
			if hand[i].Name == "Ace" && hand[i].Value == 11 {
				acesAtEleven = append(acesAtEleven, &hand[i])
			}
		}
		if(len(acesAtEleven) > 0){
			for _, cardPtr := range acesAtEleven {
				cardPtr.Value = 1
				fmt.Println(hand)
				value = DeterminehandValueSimple(hand)
				if value <= 21 {
					break
				}
			}
		}
	}
	return initPac.HandWithValue{Hand: hand, Value: value}
}

//Determines the value of the hand NOT accounting for aces yet
func DeterminehandValueSimple(hand []initPac.Card) int {
	value := 0
	for _, card := range hand {
		value += card.Value
	}
	return value
}

func DetermineIfHandIsSoft(hand []initPac.Card) bool {
	for i := range hand {
		if hand[i].Name == "Ace" && hand[i].Value == 11 {
			return true
		}
	}
	return false
}