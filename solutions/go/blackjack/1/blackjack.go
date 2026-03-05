package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":		return 11
        case "two":		return 2
        case "three":	return 3
        case "four":	return 4
    	case "five":	return 5
    	case "six":		return 6
    	case "seven":	return 7
    	case "eight":	return 8
    	case "nine":	return 9
    	case "ten", "jack", "queen", "king": return 10
        default:		return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	switch {
    	// If you have a pair of aces, you must split (P)
    	case card1 == "ace" && card2 == "ace": return "P"
    
    	// Blackjack: Win (W) if dealer doesn't have an Ace or 10-value card
    	case playerScore == 21:
    		if dealerScore < 10 {
    			return "W"
    		}
    		return "S"
    
    	// Stand (S)
    	case playerScore >= 17 && playerScore <= 20: return "S"
    	case playerScore >= 12 && playerScore <= 16 && dealerScore < 7: return "S"
    
    	// Hit (H)
    	default: return "H"
	}
}