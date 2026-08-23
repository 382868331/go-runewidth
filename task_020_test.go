package runewidth

import (
	"os"
	"testing"
)

func TestTask020EnvironmentEnablesEastAsianWidth(t *testing.T) {
	old := os.Getenv("RUNEWIDTH_EASTASIAN")
	defer os.Setenv("RUNEWIDTH_EASTASIAN", old)
	for _, v := range []string{"1", "0"} {
		os.Setenv("RUNEWIDTH_EASTASIAN", v)
		handleEnv()
		if EastAsianWidth != (v == "1") {
			t.Fatalf("value=%q enabled=%v", v, EastAsianWidth)
		}
	}
}
