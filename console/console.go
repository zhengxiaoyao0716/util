// Package console provide functions like javascript `console` object.
package console

import (
	"bufio"
	"container/list"
	"io"
	"os"
	"os/signal"
	"strings"
	"unicode"

	"time"

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

var stack = list.New()
var reader = bufio.NewReader(os.Stdin)

// ReadLine read a line of input from standart input. By default, it will output a tip symbol
// like `> ` and wait for input. you can override this behavour by the tips param. While, if
// you don't want to show anything neither default tip nor your custom, please use nil as param.
func ReadLine(tips ...interface{}) string {
	if stack.Len() == 0 {
		if len(tips) == 0 {
			cout.Print("> ")
		} else {
			if tips[0] != nil {
				cout.Println(tips...)
			}
		}

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return ReadLine(nil)
		}
		AbortInterrupt()
		PushLine(strings.TrimRightFunc(line, unicode.IsSpace))
	}
	line := stack.Front()
	if line == nil {
		return ReadLine()
	}
	return stack.Remove(line).(string)
}

// PushLine push a line of string into pre-read list, used to insert data manually.
func PushLine(line string) { stack.PushBack(line) }

var interrupt = false
var sigChan = make(chan os.Signal, 1)

// CatchInterrupt make a listen on `SIGINT` signal, and will finish the progress after twice trigger.
// You can use `TriggerInterrupt` to send signal manually, ot `AbortInterrupt` to reset count.
func CatchInterrupt() {
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		for {
			_ = <-sigChan
			if interrupt {
				os.Exit(0)
			} else {
				cout.Println("\n(Press", cout.Info("^C"), "again to exit.)")
				interrupt = true
			}
		}
	}()
}

// TriggerInterrupt send a `SIGINT` signal manually.
func TriggerInterrupt() {
	sigChan <- os.Interrupt
	for !interrupt {
		time.Sleep(10 * time.Millisecond)
	}
}

// AbortInterrupt reset interrupt count to prevent exit.
func AbortInterrupt() { interrupt = false }
