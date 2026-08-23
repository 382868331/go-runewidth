package runewidth

import "testing"

func TestTask002ControlWidth(t *testing.T) {
	for _, s := range []string{"\n", "\t"} {
		if got := NewCondition().StringWidth(s); got != 0 {
			t.Fatalf("%q width=%d want=0", s, got)
		}
	}
}
