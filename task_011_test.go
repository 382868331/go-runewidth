package runewidth

import "testing"

func TestTask011TruncateLeftFullyConsumed(t *testing.T) {
	for _, tc := range []struct {
		s      string
		w      int
		prefix string
	}{{"ab", 3, "<"}, {"界", 2, "…"}} {
		if got := NewCondition().TruncateLeft(tc.s, tc.w, tc.prefix); got != tc.prefix {
			t.Fatalf("got=%q want=%q", got, tc.prefix)
		}
	}
}
