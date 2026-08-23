package runewidth

import "testing"

func TestTask013WrapResetsAfterNewline(t *testing.T) {
	for _, s := range []string{"ab\ncd", "界\nab"} {
		if got := NewCondition().Wrap(s, 2); got != s {
			t.Fatalf("got=%q want=%q", got, s)
		}
	}
}
