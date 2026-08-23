package runewidth
import "testing"
func TestTask002ControlWidth(t *testing.T){if got:=NewCondition().StringWidth("\n");got!=0{t.Fatalf("width=%d want=0",got)}}
