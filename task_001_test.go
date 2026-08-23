package runewidth

import "testing"

func TestTask001ASCIIWidth(t *testing.T) {
	for _, s := range []string{"A", "~"} {
		if got := NewCondition().StringWidth(s); got != 1 {
			t.Fatalf("%q width=%d want=1", s, got)
		}
	}
}
