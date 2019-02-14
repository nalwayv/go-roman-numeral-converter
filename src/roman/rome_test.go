package roman

import (
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, val := range a {
		if val != b[idx] {
			return false
		}
	}

	return true
}

func TestWholeNumber(t *testing.T) {
	num := 1954
	res := wholeNumbers(num)

	expect := []int{1000, 900, 50, 4}

	if !equal(res, expect) {
		t.Error("failed to to whole a number")
	}
}

func TestDisplay(t *testing.T) {

	result1 := display(3, 1, 10, "V", "I", "X")
	expect1 := "III"

	result2 := display(30, 10, 90, "L", "X", "C")
	expect2 := "XXX"

	result3 := display(300, 100, 900, "D", "C", "M")
	expect3 := "CCC"

	if result1 != expect1 {
		t.Error("failed to display 1")
	}
	if result2 != expect2 {
		t.Error("failed to display 2")
	}
	if result3 != expect3 {
		t.Error("failed to display 3")
	}
}
