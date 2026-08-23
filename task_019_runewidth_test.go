package runewidth

import "testing"

// TestTask019 isolates the 空rune集 regression.
func TestTask019(t *testing.T) {
	TestTruncateNoNeeded(t)
}
