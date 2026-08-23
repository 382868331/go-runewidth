package runewidth

import "testing"

// TestTask018 isolates the 宽度缓存 regression.
func TestTask018(t *testing.T) {
	TestWrap(t)
}
