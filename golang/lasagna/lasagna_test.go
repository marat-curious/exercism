package lasagna

import "testing"

func TestRemainingOvenTime(t *testing.T) {
	var input = [4]int{ 30, -10, 0, 50 }
	var expected = [4]int{ 10, 0, 40, 0 }
	for i, v := range input {
		r := RemainingOvenTime(v)
		if r != expected[i] {
			t.Errorf("RemainingOvenTime(%d) = %d; want %d", v, r, expected[i])
		}
	}
}

func TestPreparationTime(t *testing.T) {
	var input = [2]int{ 0, 2 }
	var expected = [2]int{ 0, 4 }
	for i, v := range input {
		r := PreparationTime(v)
		if r != expected[i] {
			t.Errorf("PreparationTime(%d) = %d; want %d", v, r, expected[i])
		}
	}
}

func TestElapsedTime(t *testing.T) {
	var inputLayers = [4]int{ 2, 5, 1, 4 }
	var inputActualMinutes = [4]int{ 10, 25, 30, 8 }
	var expected = [4]int{ 14, 35, 32, 16 }
	for i, v := range inputLayers {
		r := ElapsedTime(v, inputActualMinutes[i])
		if r != expected[i] {
			t.Errorf("ElapsedTime(%d, %d) = %d; want %d", v, inputActualMinutes[i], r, expected[i])
		}
	}
}
