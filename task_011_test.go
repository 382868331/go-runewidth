package runewidth
import "testing"
func TestTask011TruncateLeftFullyConsumed(t *testing.T){if got:=NewCondition().TruncateLeft("ab",3,"<");got!="<"{t.Fatalf("got=%q want=<",got)}}
