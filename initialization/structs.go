package initialization

type Game struct {
	DealerHitSoftSeventeen bool;
	SurrenderAllowed bool;
	BjPayoutRatio float64;
	DoubleAfterSplit bool;
	BetMinimum int;
	RollingCount int;
	TrueCount int;
	Shoe CardStack;
	DealersHand []Card;
	Upcard Card;
	ShuffleAt float64;
}

type Player struct {
	Bankroll int;
	Hand []Card;
	Bet int;
	IsCounter bool;
}

type Card struct {
	Name string;
	Value int;
	Suit string;
}

type CardStack struct {
	Cards []Card;
}

type BetSpread struct {
	MinusFive int;
	MinusFour int;
	MinusThree int;
	MinusTwo int;
	MinusOne int;
	Zero int;
	One int;
	Two int;
	Three int;
	Four int;
	Five int;
}

type HandWithValue struct {
	Hand []Card;
	Value int;
}
