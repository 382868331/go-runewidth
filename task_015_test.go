package runewidth

import "testing"

func TestTask015FillLeftPadsBeforeText(t *testing.T) {
	for _, tc := range []struct {
		s    string
		w    int
		want string
	}{{"ab", 4, "  ab"}, {"界", 4, "  界"}} {
		if got := NewCondition().FillLeft(tc.s, tc.w); got != tc.want {
			t.Fatalf("got=%q want=%q", got, tc.want)
		}
	}
}
