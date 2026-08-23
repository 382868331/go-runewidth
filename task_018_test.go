package runewidth
import "testing"
func TestTask018WidthTableUpperBoundary(t *testing.T){if w,ok:=inWidthTable(21,widthTable{{10,20,2}});ok{t.Fatalf("out-of-range width=%d reported present",w)}}
