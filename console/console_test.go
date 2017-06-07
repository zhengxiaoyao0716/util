package console

import (
	"testing"

	"github.com/fatih/color"
)

func TestAll(t *testing.T) {
	Log("underline")
	Info("info")
	Warn("Warn")
	Yea("yea")
	Err("err")
	Custom(color.BgWhite, color.FgBlack)("reverse")
}
