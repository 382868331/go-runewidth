package runewidth
import "testing"
func TestTask013WrapResetsAfterNewline(t *testing.T){if got:=NewCondition().Wrap("ab\ncd",2);got!="ab\ncd"{t.Fatalf("got=%q",got)}}
