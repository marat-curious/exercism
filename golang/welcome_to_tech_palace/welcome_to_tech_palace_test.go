package techpalace

import (
	"strings"
	"testing"
)

func TestWelcomeMessage(t *testing.T) {
	const prefix string = "Welcome to the Tech Palace,"
	var inputCustomers = [3]string{ "", "JUDY", "TOM" }
	for i := 0; i < 3; i++ {
		r := WelcomeMessage(inputCustomers[i])
		e := prefix + " " + inputCustomers[i]
		if strings.Compare(r, e) != 0 {
			t.Errorf("WelcomeMessage(%q) = %q, wants %q", inputCustomers[i], r, e)
		}
	}
}

func TestAddBorder(t *testing.T) {
	var inputMessages = [3]string{ "", "Welcome!", "Judy" }
	var inputStarsNum = [3]int{ 0, 5, 10 }
	for i := 0; i < 3; i++ {
		r := AddBorder(inputMessages[i], inputStarsNum[i])
		s := strings.Repeat("*", inputStarsNum[i])
		e := s + "\n" + inputMessages[i] + "\n" + s
		if strings.Compare(r, e) != 0 {
			t.Errorf("AddBorder(%q, %d) = %q, wants %q", inputMessages[i], inputStarsNum[i], r, e)
		}
	}
}

func TestCleanupMessage(t *testing.T) {
	var inputMessages = [3]string{ "***\n**\n***", "*****\n*   Test  *\n*****", "*****\n*Another Test*\n*****" }
	var expected = [3]string{ "", "Test", "Another Test"}
	for i := 0; i < 3; i++ {
		r := CleanupMessage(inputMessages[i])
		if strings.Compare(r, expected[i]) != 0 {
			t.Errorf("CleanupMessage(%q) = %q, wants %q", inputMessages[i], r, expected[i])
		}
	}
}
