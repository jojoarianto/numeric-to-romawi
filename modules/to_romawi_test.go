package modules

import (
	"testing"
)

func TestAngkaToRomawi(t *testing.T) {

	var romawiTest = []struct {
		n        int    // input
		expected string // expected result
	}{
		{1, "I"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{10, "X"},
		{11, "XI"},
		{12, "XII"},
		{13, "XIII"},
		{14, "XIV"},
		{15, "XV"},
		{19, "XIX"},
		{20, "XX"},
		{30, "XXX"},
		{1637, "MDCXXXVII"},
	}

	for _, p := range romawiTest {
		actual := AngkaToRomawi(p.n)
		if actual != p.expected {
			t.Errorf("AngkaToRomawi(%d): expected %s, actual %s", p.n, p.expected, actual)
		}
	}

}
