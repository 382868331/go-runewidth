package runewidth
import "testing"
func TestTask005CombiningMarkWidth(t *testing.T){c:=&Condition{};if got:=c.RuneWidth('\u0301');got!=0{t.Fatalf("width=%d want=0",got)}}
