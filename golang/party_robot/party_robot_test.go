package partyrobot

import (
	"fmt"
	"strings"
	"testing"
)

func TestWelcome(t *testing.T) {
	const prefix string = "Welcome to my party,"
	var inputNames = [3]string{ "", "Christiane", "Top" }
	for i := 0; i < 3; i++ {
		r := Welcome(inputNames[i])
		e := prefix + " " + inputNames[i] + "!"
		if strings.Compare(r, e) != 0 {
			t.Errorf("Welcome(%q) = %q, wants %q", inputNames[i], r, e)
		}
	}
}

func TestHappyBirthday(t *testing.T) {
	const prefix_1 string = "Happy birthday"
	const prefix_2 string = "You are now"
	const suffix_2 string = "years old!"
	var inputNames = [3]string{ "", "Frank", "123" }
	var inputAges = [3]int{ 0, 58, 345 }
	for i := 0; i < 3; i++ {
		r := HappyBirthday(inputNames[i], inputAges[i])
		e := prefix_1 + " " + inputNames[i] + "! " + prefix_2 + " " + fmt.Sprint(inputAges[i]) + " " + suffix_2
		if strings.Compare(r, e) != 0 {
			t.Errorf("HappyBirthday(%q, %d) = %q, wants %q", inputNames[i], inputAges[i], r, e)
		}
	}
}

func TestAssignTable(t *testing.T) {
	const s1 string = "Welcome to my party,"
	const s2 string = "You have been assigned to table"
	const s3 string = "Your table is"
	const s4 string = "exactly"
	const s5 string = "meters from here."
	const s6 string = "You will be sitting next to"
	var inputNames = [3]string{ "Christiane", "Frank", "Tom" }
	var inputTables = [3]int{ 1, 10, 85 }
	var inputNeighbors = [3]string{ "Frank", "Tom", "Christiane" }
	var inputDirections = [3]string{ "on the left", "on the right", "on the left" }
	var inputDistance = [3]float64{ 23.7834298, 10.5748, 54.0 }
	for i := 0; i < 3; i++ {
		r := AssignTable(inputNames[i], inputTables[i], inputNeighbors[i], inputDirections[i], inputDistance[i])
		e := s1 + " " + inputNames[i] + "!\n" + s2 + " " + fmt.Sprintf("%03d", inputTables[i]) + ". " + s3 + " " + inputDirections[i] + ", " + s4 + " " + fmt.Sprintf("%2.2f", inputDistance[i]) + " " + s5 + "\n" + s6 + " " + inputNeighbors[i] + "."
		if strings.Compare(r, e) != 0 {
			t.Errorf("AssignTable(%q, %d, %q, %q, %f) = %q, wants %q", inputNames[i], inputTables[i], inputNeighbors[i], inputDirections[i], inputDistance[i], r, e)
		}
	}
}
