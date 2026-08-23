package runewidth

import "testing"

func TestTask018WidthTableUpperBoundary(t *testing.T) {
	tab := widthTable{{10, 20, 2}}
	for _, r := range []rune{21, 99} {
		if w, ok := inWidthTable(r, tab); ok {
			t.Fatalf("rune=%d width=%d reported present", r, w)
		}
	}
}
