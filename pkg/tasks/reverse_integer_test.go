package tasks

import (
	"testing"
)

func TestReverseInteger(t *testing.T) {
	num := 123
	result := Reverse(num)
	if result != 321 {
		t.Fatalf("num: 123, expected 321, result: %d", result)
	}

	num = 210
	if Reverse(num) != 12 {
		t.Fatalf("num: 210, expected 12, result: %d", result)
	}

	num = -120
	if Reverse(num) != -21 {
		t.Fatalf("num: -120, expected -21, result: %d", result)
	}

}
