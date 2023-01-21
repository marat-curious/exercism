package blackjack

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
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {
	var c1 int = ParseCard(card1)
	var c2 int = ParseCard(card2)
	var d int = ParseCard(dealerCard)
	var s int = c1 + c2

	switch s {
	case 22:
		return "P"
	case 21:
		if d < 10 {
			return "W"
		} else {
			return "S"
		}
	case 20, 19, 18, 17:
		return "S"
	case 16, 15, 14, 13, 12:
		if d >= 7 {
			return "H"
		} else {
			return "S"
		}
	default:
		return "H"
	}
}
