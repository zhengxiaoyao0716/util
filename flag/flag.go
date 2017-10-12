// Package flag used to extends official `flag` package.
// It can give accurate info that if user really assign the value.
// Notice that I only add the methods I need in practice,
// so that if need more, maybe you have to add them manual.
package flag

import "strconv"
import "flag"
import "fmt"

// string

type stringValue func() string

func (s *stringValue) Get() interface{} {
	if *s == nil {
		return ""
	}
	return (*s)()
}

func (s *stringValue) String() string { return s.Get().(string) }

func (s *stringValue) Set(val string) error {
	*s = func() string { return val }
	return nil
}

// int
type intValue func() int

func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = func() int { return int(v) }
	return err
}

func (i *intValue) Get() interface{} {
	if *i == nil {
		return 0
	}
	return (*i)()
}

func (i *intValue) String() string { return strconv.Itoa(i.Get().(int)) }

// float
type floatValue func() float64

func (f *floatValue) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = func() float64 { return v }
	return err
}

func (f *floatValue) Get() interface{} {
	if *f == nil {
		return 0
	}
	return (*f)()
}

func (f *floatValue) String() string { return fmt.Sprint(f.Get().(float64)) }

// bool
type boolValue func() bool

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = func() bool { return v }
	return err
}

func (b *boolValue) Get() interface{} {
	if *b == nil {
		return false
	}
	return (*b)()
}

func (b *boolValue) String() string { return strconv.FormatBool(b.Get().(bool)) }

func (b *boolValue) IsBoolFlag() bool { return true }

// String .
func String(name string, usage string) *func() string {
	p := new(stringValue)
	flag.Var(p, name, usage)
	return (*func() string)(p)
}

// Int .
func Int(name string, usage string) *func() int {
	p := new(intValue)
	flag.Var(p, name, usage)
	return (*func() int)(p)
}

// Float .
func Float(name string, usage string) *func() float64 {
	p := new(floatValue)
	flag.Var(p, name, usage)
	return (*func() float64)(p)
}

// Bool .
func Bool(name string, usage string) *func() bool {
	p := new(boolValue)
	flag.Var((*boolValue)(p), name, usage)
	return (*func() bool)(p)
}

// Parse .
func Parse() { flag.Parse() }

// CommandLine .
var CommandLine = flag.CommandLine
