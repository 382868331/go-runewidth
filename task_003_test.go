package runewidth

import "testing"

func TestTask003DeleteWidth(t *testing.T) {
	for _, b := range []byte{0x7f, 0x1f} {
		if got := NewCondition().StringWidth(string([]byte{b})); got != 0 {
			t.Fatalf("byte=%#x width=%d", b, got)
		}
	}
}
