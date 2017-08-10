// Package cout means 'color output', for it decorated words with color.
// It can also understood as a simulator for cpp output, for this concat style:
// cout.Println(Log(), Info(), ...) ~= cout << log() << Info() << ...
package cout

import (
	"fmt"

	"github.com/fatih/color"
)

// Palette declare the color or type for different tags.
type Palette struct {
	LOG  color.Attribute
	INFO color.Attribute
	WARN color.Attribute
	YES  color.Attribute
	ERR  color.Attribute
}

var colors [5]*color.Color

// SetPalette set the color or type of each taget.
func SetPalette(p Palette) {
	colors = [5]*color.Color{}
	for index, attr := range []color.Attribute{p.LOG, p.INFO, p.WARN, p.YES, p.ERR} {
		colors[index] = color.New(attr)
	}
}

// Log wrap the give string with the color of `LOG` taget.
func Log(format string, a ...interface{}) string { return colors[0].SprintfFunc()(format, a...) }

// Info wrap the give string with the color of `INFO` taget.
func Info(format string, a ...interface{}) string { return colors[1].SprintfFunc()(format, a...) }

// Warn wrap the give string with the color of `WARN` taget.
func Warn(format string, a ...interface{}) string { return colors[2].SprintfFunc()(format, a...) }

// Yes wrap the give string with the color of `YES` taget.
func Yes(format string, a ...interface{}) string { return colors[3].SprintfFunc()(format, a...) }

// Err wrap the give string with the color of `ERR` taget.
func Err(format string, a ...interface{}) string { return colors[4].SprintfFunc()(format, a...) }

// Custom create a custom decorate function to wrap the string.
func Custom(attrs ...color.Attribute) func(format string, a ...interface{}) string {
	c := color.New(attrs...)
	return func(format string, a ...interface{}) string { return c.SprintfFunc()(format, a...) }
}

// Printf print the strings with format.
func Printf(format string, a ...interface{}) { fmt.Fprintf(color.Output, format, a...) }

// Print print the strings without new line.
func Print(a ...interface{}) { fmt.Fprint(color.Output, a...) }

// Println print the strings in line.
func Println(a ...interface{}) { fmt.Fprintln(color.Output, a...) }

func init() {
	SetPalette(Palette{
		LOG:  color.FgHiWhite,
		INFO: color.FgHiBlue,
		WARN: color.FgHiYellow,
		YES:  color.FgHiGreen,
		ERR:  color.FgHiRed,
	})
}
