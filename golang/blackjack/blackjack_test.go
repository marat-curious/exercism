package blackjack

import (
	"strings"
	"testing"
)

func TestParseCard(t *testing.T) {
	var cards = [13]string{ "ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king" }
	var values = [13]int{ 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10 }
	for i := 0; i < 13; i++ {
		r := ParseCard(cards[i])
		if r != values[i] {
			t.Errorf("ParseCard(%q) = %d, wants %d", cards[i], r, values[i])
		}
	}
}

func TestFirstTurn(t *testing.T) {
	var cards1 = [5]string{ "ace", "ace", "five", "two", "ten" }
	var cards2 = [5]string{ "ace", "king", "queen", "six", "nine" }
	var dealerCards = [5]string{ "jack", "ace", "ace", "ten", "jack" }
	var expected = [5]string{ "P", "S", "H", "H", "S" }
	for i := 0; i < 5; i++ {
		r := FirstTurn(cards1[i], cards2[i], dealerCards[i])
		if strings.ToUpper(r) != expected[i] {
			t.Errorf("FirstTurn(%q, %q, %q) = %q, wants %q", cards1[i], cards2[i], dealerCards[i], r, expected[i])
		}
	}
}
