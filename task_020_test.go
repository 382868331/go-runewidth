package runewidth
import("os";"testing")
func TestTask020EnvironmentEnablesEastAsianWidth(t *testing.T){old:=os.Getenv("RUNEWIDTH_EASTASIAN");defer os.Setenv("RUNEWIDTH_EASTASIAN",old);os.Setenv("RUNEWIDTH_EASTASIAN","1");handleEnv();if !EastAsianWidth{t.Fatal("environment flag did not enable East Asian width")}}
