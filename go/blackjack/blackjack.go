package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
/*
Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

If you have a pair of aces you must always split them.
If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win.
If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
If your cards sum up to a value within the range [17, 20] you should always stand.
If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
If your cards sum up to 11 or lower you should always hit.
*/
func FirstTurn(card1, card2, dealerCard string) string {
	count := ParseCard(card1) + ParseCard(card2)

	switch {
	case count == 22:
		return "P"
	case count == 21:
		if ParseCard(dealerCard) > 9 {
			return "S"
		} else {
			return "W"
		}
	case count > 16 && count < 21:
		return "S"
	case count > 11 && count < 17:
		if ParseCard(dealerCard) > 6 {
			return "H"
		} else {
			return "S"
		}
	default:
		return "H"
	}
}
