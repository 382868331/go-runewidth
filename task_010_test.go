package runewidth

import "testing"

func TestTask010TruncateStopsBeforeWideCluster(t *testing.T) {
	for _, tc := range []struct {
		s    string
		w    int
		want string
	}{{"A界B", 2, "A"}, {"界AB", 1, ""}} {
		if got := NewCondition().Truncate(tc.s, tc.w, ""); got != tc.want {
			t.Fatalf("got=%q want=%q", got, tc.want)
		}
	}
}
