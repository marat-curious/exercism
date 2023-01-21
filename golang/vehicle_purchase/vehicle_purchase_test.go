package purchase

import (
	"strings"
	"testing"
)

func TestNeedsLicense(t *testing.T) {
	var inputKinds = [3]string{ "car", "bike", "truck" }
	var expected = [3]bool{ true, false, true }
	for i := 0; i < 3; i++ {
		r := NeedsLicense(inputKinds[i])
		if r != expected[i] {
			t.Errorf("NeedsLicense(%q) = %t, wants %t", inputKinds[i], r, expected[i])
		}
	}
}

func TestChooseVehicle(t *testing.T) {
	const suffix = "is clearly the better choice."
	var inputOptions1 = [2]string{ "Wuling Hongguang", "Volkswagen Beetle" }
	var inputOptions2 = [2]string{ "Toyota Corolla", "Volkswagen Golf" }
	var expected = [2]string{ "Toyota Corolla", "Volkswagen Beetle" }
	for i := 0; i < 2; i++ {
		r := ChooseVehicle(inputOptions1[i], inputOptions2[i])
		if strings.Compare(r, expected[i] + " " + suffix) != 0 {
			t.Errorf("ChooseVehicle(%q, %q) = %q, wants %q", inputOptions1[i], inputOptions2[i], r, expected[i] + " " + suffix)
		}
	}
}

func TestCalculateResellPrice(t *testing.T) {
	var inputOriginalPrices = [3]float64{ 1000.0, 1000.0, 1000.0 }
	var inputAges = [3]float64{ 1, 5, 15 }
	var expected = [3]float64{ 800.0, 700.0, 500.0 }
	for i := 0; i < 3; i++ {
		r := CalculateResellPrice(inputOriginalPrices[i], inputAges[i])
		if int(r) != int(expected[i]) {
			t.Errorf("CalculateResellPrice(%f, %f) = %f, wants %f", inputOriginalPrices[i], inputAges[i], r, expected[i])
		}
	}
}
