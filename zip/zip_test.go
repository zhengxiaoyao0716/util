package zip

import (
	"os"
	"testing"
)

func TestPack(t *testing.T) {
	dst, src := "temp.zip", "./"
	if err := Pack(dst, src); err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		t.Fatal("Packed file doesn't found.")
	}
	if err := os.Remove(dst); err != nil {
		t.Error(err)
	}
}
