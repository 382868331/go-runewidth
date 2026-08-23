package runewidth
import "testing"
func TestTask007EmojiClusterCap(t *testing.T){for _,s:=range []string{"👨‍👩‍👧","🇨🇳"}{if got:=NewCondition().StringWidth(s);got!=2{t.Fatalf("%q width=%d",s,got)}}}
