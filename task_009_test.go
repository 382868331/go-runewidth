package runewidth

import "testing"

func TestTask009TruncateReservesTail(t *testing.T) {
	for _, tc := range []struct {
		s          string
		w          int
		tail, want string
	}{{"abcdef", 4, "..", "ab.."}, {"界abc", 4, "…", "界a…"}} {
		if got := NewCondition().Truncate(tc.s, tc.w, tc.tail); got != tc.want {
			t.Fatalf("got=%q want=%q", got, tc.want)
		}
	}
}
