package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value := 0
	switch {
	case card == "ace":
		value = 11
	case card == "two":
		value = 2
	case card == "three":
		value = 3
	case card == "four":
		value = 4
	case card == "five":
		value = 5
	case card == "six":
		value = 6
	case card == "seven":
		value = 7
	case card == "eight":
		value = 8
	case card == "nine":
		value = 9
	case card == "ten":
		value = 10
	case card == "jack":
		value = 10
	case card == "queen":
		value = 10
	case card == "king":
		value = 10
	default:
		value = 0
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	action := ""
	card1Int := ParseCard(card1)
	card2Int := ParseCard(card2)
	dealerCardInt := ParseCard(dealerCard)
	cardSum := card1Int + card2Int

	switch {
	case card1Int == 11 && card2Int == 11:
		action = "P"
	case cardSum >= 17 && cardSum <= 20:
		action = "S"
	case cardSum >= 12 && cardSum <= 16 && dealerCardInt < 7:
		action = "S"
	case cardSum >= 12 && cardSum <= 16 && dealerCardInt >= 7:
		action = "H"
	case cardSum <= 11:
		action = "H"
	case cardSum == 21 && dealerCardInt < 10:
		action = "W"
	case cardSum == 21 && dealerCardInt >= 10:
		action = "S"
	}

	return action
}
