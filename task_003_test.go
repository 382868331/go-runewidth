package runewidth
import "testing"
func TestTask003DeleteWidth(t *testing.T){if got:=NewCondition().StringWidth(string([]byte{0x7f}));got!=0{t.Fatalf("width=%d want=0",got)}}
