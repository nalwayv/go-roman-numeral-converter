package roman

import (
	"fmt"
	"math"
	"strings"
)

// wholeNumbers ... convert a number to its 1, 10, 100, 1000 ... components
// 1024 -> [1000, 20, 4]
// 512  -> [50, 10, 2]
// 8    -> [8]
func wholeNumbers(num int) []int {

	roundToTen := func(num int) int {
		if num <= 0 {
			return 0
		}
		val := math.Pow(10.0, math.Ceil(math.Log10(float64(num))))
		return int(val / 10)
	}

	div := roundToTen(num)
	res := make([]int, 0)

	for num > 0 {
		v := num / div
		v *= div
		if v != 0 {
			res = append(res, v)
		}
		num -= v
		div /= 10
	}

	return res
}

// display ...
// https://www.mathsisfun.com/roman-numerals.html
// When a symbol appears after a larger (or equal) symbol it is added
// 		Example: VI = V + I = 5 + 1 = 6
// 		Example: LXX = L + X + X = 50 + 10 + 10 = 70

// But if the symbol appears before a larger symbol it is subtracted
// 		Example: IV = V − I = 5 − 1 = 4
// 		Example: IX = X − I = 10 − 1 = 9
func display(num, scale, max int, sym1, sym2, sym3 string) string {

	result := make([]string, 0)

	// mid
	if num == 5*scale {
		result = append(result, sym1)
		return strings.Join(result, "")
	}

	// less or more than mid
	if num < (max / 2) {
		if num == 4*scale {
			result = append(result, fmt.Sprintf("%s%s", sym2, sym1))
		} else {
			numberOfSymbols := make([]string, 0)
			for i := 0; i < (num / scale); i++ {
				numberOfSymbols = append(numberOfSymbols, sym2)
			}
			result = append(result, strings.Join(numberOfSymbols, ""))
		}
	}

	// more than mid
	if num > (max / 2) {
		if num == 9*scale {
			result = append(result, fmt.Sprintf("%s%s", sym2, sym3))
		} else {

			numberOfSymbols := make([]string, 0)
			numberOfSymbols = append(numberOfSymbols, sym1)

			for i := 0; i < num/scale-5; i++ {
				numberOfSymbols = append(numberOfSymbols, sym2)
			}

			result = append(result, strings.Join(numberOfSymbols, ""))
		}
	}

	return strings.Join(result, "")
}

// ToRoman ... convert an int number to a roman number
func ToRoman(value int) string {
	roman := make([]string, 0)

	for _, num := range wholeNumbers(value) {
		if num < 10 {
			roman = append(roman, display(num, 1, 10, "V", "I", "X"))
		} else if num >= 10 && num < 91 {
			roman = append(roman, display(num, 10, 90, "L", "X", "C"))
		} else if num >= 100 && num < 901 {
			roman = append(roman, display(num, 100, 900, "D", "C", "M"))
		} else {
			numberOfSymbols := make([]string, 0)
			for i := 0; i < (num / 1000); i++ {
				numberOfSymbols = append(numberOfSymbols, "M")
			}
			roman = append(roman, strings.Join(numberOfSymbols, ""))
		}
	}

	return strings.Join(roman, "")
}
