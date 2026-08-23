package runewidth
import "testing"
func TestTask009TruncateReservesTail(t *testing.T){if got:=NewCondition().Truncate("abcdef",4,"..");got!="ab.."{t.Fatalf("got=%q want=ab..",got)}}
