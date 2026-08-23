package runewidth
import "testing"
func TestTask012PrefixConsumesBudget(t *testing.T){if got:=NewCondition().TruncatePrefix("abcdef",2,"..");got!=".."{t.Fatalf("got=%q want=..",got)}}
