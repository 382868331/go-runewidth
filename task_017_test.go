package runewidth
import "testing"
func TestTask017InTableLowerBoundary(t *testing.T){if inTable(9,table{{10,20}}){t.Fatal("rune below table reported present")}}
