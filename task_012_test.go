package runewidth

import "testing"

func TestTask012PrefixConsumesBudget(t *testing.T) {
	for _, tc := range []struct {
		s string
		w int
		p string
	}{{"abcdef", 2, ".."}, {"界abc", 2, "界"}} {
		if got := NewCondition().TruncatePrefix(tc.s, tc.w, tc.p); got != tc.p {
			t.Fatalf("got=%q want=%q", got, tc.p)
		}
	}
}
