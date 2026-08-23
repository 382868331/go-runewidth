package runewidth

import "testing"

func TestTask017InTableLowerBoundary(t *testing.T) {
	tab := table{{10, 20}}
	for _, r := range []rune{0, 9} {
		if inTable(r, tab) {
			t.Fatalf("rune %d below table reported present", r)
		}
	}
}
