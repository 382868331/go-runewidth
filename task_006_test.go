package runewidth
import "testing"
func TestTask006WideRuneWidth(t *testing.T){c:=&Condition{};if got:=c.RuneWidth('界');got!=2{t.Fatalf("width=%d want=2",got)}}
