package runewidth
import "testing"
func TestTask001ASCIIWidth(t *testing.T){if got:=NewCondition().StringWidth("A");got!=1{t.Fatalf("width=%d want=1",got)}}
