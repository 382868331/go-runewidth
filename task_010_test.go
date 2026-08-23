package runewidth
import "testing"
func TestTask010TruncateStopsBeforeWideCluster(t *testing.T){if got:=NewCondition().Truncate("A界B",2,"");got!="A"{t.Fatalf("got=%q want=A",got)}}
