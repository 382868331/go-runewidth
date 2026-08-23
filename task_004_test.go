package runewidth
import "testing"
func TestTask004InvalidRuneWidth(t *testing.T){if got:=NewCondition().RuneWidth(-1);got!=0{t.Fatalf("width=%d want=0",got)}}
