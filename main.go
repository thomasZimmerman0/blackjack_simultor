package main

import(
	"fmt"
	initPac "github.com/thomasZimmerman0/blackjack_simultor/initialization"
	playHand "github.com/thomasZimmerman0/blackjack_simultor/play_hand"
)

func main(){
	// shuffledShoe := initPac.BuildDecks(6);

	// count := 0
	// for _, card := range shuffledShoe.Cards {
	// 	count++
	// 	fmt.Println(card)
	// }

	// game := initPac.Game{
	// 	DealerHitSoftSeventeen: true,
	// 	SurrenderAllowed: false,
	// 	BjPayoutRatio: 3 / 2,
	// 	DoubleAfterSplit: false,
	// 	RollingCount: 0,
	// 	TrueCount: 0,
	// 	Shoe: shuffledShoe,
	// 	DealersHand: []initPac.Card{},
	// 	Upcard: initPac.Card{Name: "none", Value: 0, Suit: "none"},
	// }

	player := initPac.Player{
		Bankroll: 500,
		Hand: []initPac.Card{
			{Name: "Six", Value: 6, Suit: "Spades"},
			{Name: "King", Value: 10, Suit: "Spades"},
			{Name: "Jack", Value: 10, Suit: "Spades"},
		},
		Bet: 15,
		IsCounter: true,
	}

	fmt.Println(playHand.DetermineHandValueInFull(player.Hand))

	// dealtCard := playHand.DealCard(&game)

	// for _, card := range game.Shoe.Cards {
	// 	count++
	// 	fmt.Println(card)
	// }

	
	// fmt.Println();
	// fmt.Println(count)
	// fmt.Println(dealtCard)
	// fmt.Println();
	// count = 0

	// dealtCard = playHand.DealCard(&game)

	// for _, card := range game.Shoe.Cards {
	// 	count++
	// 	fmt.Println(card)
	// }

	
	// fmt.Println();
	// fmt.Println(count)
	// fmt.Println(dealtCard)
	// fmt.Println();
	// count = 0

}

