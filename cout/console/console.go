package console

import (
	"github.com/fatih/color"
	"github.com/zhengxiaoyao0716/util/cout"
)

// Log wrap the give string with the color of `LOG` taget.
func Log(format string, a ...interface{}) { cout.Println(cout.Log(format, a...)) }

// Info wrap the give string with the color of `INFO` taget.
func Info(format string, a ...interface{}) { cout.Println(cout.Info(format, a...)) }

// Warn wrap the give string with the color of `WARN` taget.
func Warn(format string, a ...interface{}) { cout.Println(cout.Warn(format, a...)) }

// Yea wrap the give string with the color of `YEA` taget.
func Yea(format string, a ...interface{}) { cout.Println(cout.Yea(format, a...)) }

// Err wrap the give string with the color of `ERR` taget.
func Err(format string, a ...interface{}) { cout.Println(cout.Err(format, a...)) }

// Custom create a custom decorate function to wrap the string.
func Custom(attrs ...color.Attribute) func(format string, a ...interface{}) {
	c := color.New(attrs...)
	return func(format string, a ...interface{}) { cout.Println(c.SprintfFunc()(format, a...)) }
}
