package runewidth

import "testing"

func TestTask005CombiningMarkWidth(t *testing.T) {
	c := &Condition{}
	for _, r := range []rune{'\u0301', '\u200d'} {
		if got := c.RuneWidth(r); got != 0 {
			t.Fatalf("rune=%U width=%d", r, got)
		}
	}
}
