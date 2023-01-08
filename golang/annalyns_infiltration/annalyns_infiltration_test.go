package annalyn

import (
	"strconv"
	"testing"
)

func generateInputParams() [8][3]bool {
	result := [8][3]bool{}
	for i := 0; i < 8; i++ {
		input := [3]bool{ false, false, false }
		r := strconv.FormatInt(int64(i), 2)
		for j, v := range r {
			if diff := 0; v == 49 { // compare to digit "1" (49 - unicode encoding)
				if len(r) < 3 {
					diff = 3 - len(r)
				}
				input[j + diff] = true
			}
		}
		result[i] = input
	}
	return result
}

func TestCanFastAttack(t *testing.T) {
	input := [2]bool{ true, false }
	expected := [2]bool{ false, true }
	for i := 0; i < 2; i++ {
		r := CanFastAttack(input[i])
		if r != expected[i] {
			t.Errorf("CanFastAttack(%t) = %t, want %t", input[i], r, expected[i])
		}
	}
}

func TestCanSpy(t *testing.T) {
	input := generateInputParams()
	expected := [8]bool{ false, true, true, true, true, true, true, true }
	for i, v := range input {
		p := CanSpy(v[0], v[1], v[2])
		if p != expected[i] {
			t.Errorf("CanSpy(%t, %t, %t) = %t, want %t", v[0], v[1], v[2], p, expected[i])
		}
	}
}

func TestCanSignalPrisoner(t *testing.T) {
	input := [4][2]bool{ { false, false }, { false, true }, { true, false }, { true, true } }
	expected := [4]bool{ false, true, false, false }
	for i := 0; i < 4; i++ {
		r := CanSignalPrisoner(input[i][0], input[i][1])
		if r != expected[i] {
			t.Errorf("CanSignalPrisoner(%t, %t) = %t, want %t", input[i][0], input[i][1], r, expected[i])
		}
	}
}
