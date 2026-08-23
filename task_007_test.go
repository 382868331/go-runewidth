package runewidth
import "testing"
func TestTask007EmojiClusterCap(t *testing.T){if got:=NewCondition().StringWidth("👨‍👩‍👧");got!=2{t.Fatalf("width=%d want=2",got)}}
