package runewidth

import "testing"

func TestTask016MergeAdjacentIntervals(t *testing.T) {
	for _, tc := range []struct {
		a, b table
		want interval
	}{{table{{1, 2}}, table{{3, 4}}, interval{1, 4}}, {table{{5, 7}}, table{{7, 9}}, interval{5, 9}}} {
		got := mergeIntervals(tc.a, tc.b)
		if len(got) != 1 || got[0] != tc.want {
			t.Fatalf("merged=%v", got)
		}
	}
}
