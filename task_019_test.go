package runewidth

import "testing"

func TestTask019PackedLUTKeepsOddRuneWidth(t *testing.T) {
	c := &Condition{}
	c.CreateLUT()
	for _, r := range []rune{'A', '界'} {
		want := runeWidthNoLUT(r, false, true)
		if got := c.RuneWidth(r); got != want {
			t.Fatalf("rune=%U width=%d want=%d", r, got, want)
		}
	}
}
