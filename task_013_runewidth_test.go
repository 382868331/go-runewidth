package runewidth

import "testing"

// TestTask013 isolates the Windows换行 regression.
func TestTask013(t *testing.T) {
	TestStringWidthInvalid(t)
}
