package bat

import "testing"

func TestAll(t *testing.T) {
	Exec(
		"echo @echo off > test.bat",
		"echo echo success >> test.bat",
		"echo pause >> test.bat",
		"echo del test.bat >> test.bat",
		"start test.bat",
	)
}
