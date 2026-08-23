package runewidth

import "testing"

func TestTask004InvalidRuneWidth(t *testing.T) {
	for _, r := range []rune{-1, 0x110000} {
		if got := NewCondition().RuneWidth(r); got != 0 {
			t.Fatalf("rune=%U width=%d", r, got)
		}
	}
}
