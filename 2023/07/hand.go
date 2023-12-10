package day07

const (
	HIGH_CARD  = 0
	ONE_PAIR   = 1
	TWO_PAIR   = 2
	THREE_SAME = 3
	FULL_HOUSE = 4
	FOUR_SAME  = 5
	FIVE_SAME  = 6
)

type hand struct {
	cards []string
	value int
	bid   int
}

func newHand(cards []string, bid int) *hand {
	hand := &hand{cards: cards, bid: bid}
	return hand
}
