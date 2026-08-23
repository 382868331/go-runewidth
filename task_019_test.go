package runewidth
import "testing"
func TestTask019PackedLUTKeepsOddRuneWidth(t *testing.T){c:=&Condition{};c.CreateLUT();if got:=c.RuneWidth('A');got!=1{t.Fatalf("width=%d want=1",got)}}
