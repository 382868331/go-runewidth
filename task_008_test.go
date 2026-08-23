package runewidth

import "testing"

func TestTask008ASCIIStreamSkipsDelete(t *testing.T) {
	for _, s := range []string{"a\x7fb", "\x7fxy"} {
		if got := NewCondition().StringWidth(s); got != 2 {
			t.Fatalf("%q width=%d", s, got)
		}
	}
}
