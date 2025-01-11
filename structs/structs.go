package structs

// import ( 
// 	"fmt" 
// )

type Game struct {
	dealerHitSoftSeventeen bool;
	surrenderAllowed bool;
	bjPayoutRatio float64;
	doubleAfterSplit float64;
	betMinimum int;
}

type Player struct {
	bankroll int;
}

type Card struct {
	name string;
	value int;
	suit string;
}

type Deck struct {
	cards [52]Card;
}