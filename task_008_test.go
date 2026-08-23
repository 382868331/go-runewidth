package runewidth
import "testing"
func TestTask008ASCIIStreamSkipsDelete(t *testing.T){if got:=NewCondition().StringWidth("a\x7fb");got!=2{t.Fatalf("width=%d want=2",got)}}
