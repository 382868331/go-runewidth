package runewidth

import "testing"

// TestTask005 isolates the 宽度收尾 regression.
func TestTask005(t *testing.T) {
	TestRuneWidthChecksums(t)
}
