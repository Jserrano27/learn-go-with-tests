package romans

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Roman       string
	}{
		{"1 gets converted to I", 1, "I"},
		{"1 gets converted to I", 2, "II"},
		{"3 gets converted to I", 3, "III"},
		{"4 gets converted to IV (can't repeat letter more than three times", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"6 gets converted to VI", 6, "VI"},
		{"7 gets converted to VII", 7, "VII"},
		{"8 gets converted to VIII", 8, "VIII"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"11 gets converted to XI", 11, "XI"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
		{"89 gets converted to LXXXIX", 89, "LXXXIX"},
		{"90 gets converted to XC", 90, "XC"},
		{"98 gets converted to XCVIII", 98, "XCVIII"},
		{"100 gets converted to C", 100, "C"},
		{"387 gets converted to CCCLXXXVII", 387, "CCCLXXXVII"},
		{"400 gets converted to CD", 400, "CD"},
		{"487 gets converted to CDLXXXVII", 487, "CDLXXXVII"},
		{"500 gets converted to D", 500, "D"},
		{"864 gets converted to DCCCLXIV", 864, "DCCCLXIV"},
		{"900 gets converted to CM", 900, "CM"},
		{"1000 gets converted to M", 1000, "M"},
		{"1872 gets converted to MDCCCLXXII", 1872, "MDCCCLXXII"},
		{"3624 gets converted to MMMDCXXIV", 3624, "MMMDCXXIV"},
	}

	for _, test := range cases {

		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Roman

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
