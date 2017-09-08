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

// Yes wrap the give string with the color of `YES` taget.
func Yes(format string, a ...interface{}) { cout.Println(cout.Yes(format, a...)) }

// Err wrap the give string with the color of `ERR` taget.
func Err(format string, a ...interface{}) { cout.Println(cout.Err(format, a...)) }

// Custom create a custom decorate function to wrap the string.
func Custom(attrs ...color.Attribute) func(format string, a ...interface{}) {
	c := color.New(attrs...)
	return func(format string, a ...interface{}) { cout.Println(c.SprintfFunc()(format, a...)) }
}

var stack = list.New()
var reader = bufio.NewReader(os.Stdin)

// In is the default tip syntax for user input, used when `tips` param is nil.
var In = cout.Yes("> ")

// ReadLine read a line of input from standart input. By default, it will output a tip symbol
// like `> ` and wait for input. you can override this behavour by the tips param. While, if
// you don't want to show anything neither default tip nor your custom, please use nil as param.
func ReadLine(tips ...interface{}) string {
	line := stack.Front()
	if line == nil {
		if len(tips) == 0 {
			cout.Print(In)
		} else if tips[0] != nil {
			cout.Print(tips...)
		}

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return ReadLine(nil)
		}
		AbortInterrupt()
		return strings.TrimRightFunc(line, unicode.IsSpace)
	}
	return stack.Remove(line).(string)
}

// ReadWord split the result of `ReadLine` and return the first field.
// Others words remand would be re-push into the stack.
func ReadWord(tips ...interface{}) string {
	line := ReadLine(tips...)
	if line == "" {
		return ""
	}

	words := strings.Fields(line)
	for i := len(words) - 1; i > 0; i-- {
		stack.PushFront(words[i])
	}
	return words[0]
}

// ReadPass behavours same as `ReadLine`, but would hide the echo of user input.
func ReadPass(tips ...interface{}) string {
	result := ReadLine(append(tips, "\033[8m")...)
	cout.Print("\033[28m")
	return result
}

// PushLine push a line of string into pre-read list, used to insert data manually.
func PushLine(line string) { stack.PushBack(line) }

var interrupt = false
var sigChan = make(chan os.Signal, 1)

// CatchInterrupt make a listen on `SIGINT` signal, and will finish the progress after twice trigger.
// You can use `TriggerInterrupt` to send signal manually, ot `AbortInterrupt` to reset count.
func CatchInterrupt(clean ...func()) {
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		for {
			_ = <-sigChan
			if interrupt {
				for _, fn := range clean {
					fn()
				}
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
