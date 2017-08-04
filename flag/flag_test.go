package flag

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	s := String("str", "test string flag.")
	sn := String("str null", "test string flag.")
	i := Int("int", "test int flag.")
	in := Int("int null", "test int flag.")
	b := Bool("bool", "test bool flag.")
	bn := Bool("bool null", "test bool flag.")

	CommandLine.Parse([]string{"-str", "str", "-int", "0", "-bool"})

	fmt.Printf("%s %d %t | %p %p %p\n", (*s)(), (*i)(), (*b)(), *sn, *in, *bn)

	if *sn != nil || *in != nil || *bn != nil {
		t.Errorf("those values should be nil")
	}
}
