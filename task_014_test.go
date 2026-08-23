package runewidth
import "testing"
func TestTask014WrapAllowsExactFit(t *testing.T){if got:=NewCondition().Wrap("ab",2);got!="ab"{t.Fatalf("got=%q want=ab",got)}}
