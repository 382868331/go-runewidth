package runewidth
import "testing"
func TestTask015FillLeftPadsBeforeText(t *testing.T){if got:=NewCondition().FillLeft("ab",4);got!="  ab"{t.Fatalf("got=%q",got)}}
