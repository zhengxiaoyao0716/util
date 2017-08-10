package cout

import (
	"testing"

	"github.com/fatih/color"
)

func TestPrintln(t *testing.T) {
	Println(
		">",
		Log("underline"), Info("info"), Warn("Warn"), Yes("yes"), Err("err"),
		Custom(color.BgWhite, color.FgBlack)("reverse"),
	)
}

func TestPrintf(t *testing.T) {
	Printf(
		"> %s %s %s %s %s %s\n",
		Log("%s", "underline"),
		Info("%s", "info"),
		Warn("%s", "Warn"),
		Yes("%s", "yes"),
		Err("%s", "err"),
		Custom(color.BgWhite, color.FgBlack)("reverse"),
	)
}
