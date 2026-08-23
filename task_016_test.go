package runewidth
import "testing"
func TestTask016MergeAdjacentIntervals(t *testing.T){got:=mergeIntervals(table{{1,2}},table{{3,4}});if len(got)!=1||got[0]!=(interval{1,4}){t.Fatalf("merged=%v",got)}}
