package safefile

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// MoveAway can to move away the file instead of delete it, used to avoid the file exist conflict.
func MoveAway(fp string) error {
	if _, err := os.Lstat(fp); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		if !os.IsExist(err) {
			return err
		}
	}
	ext := ""
	for i := len(fp) - 1; i >= 0 && !os.IsPathSeparator(fp[i]); i-- {
		if fp[i] == '.' {
			ext = fp[i:]
			fp = fp[0:i]
			break
		}
	}
	err := os.Rename(fp+ext, fmt.Sprintf("%s.%d%s", fp, time.Now().Unix(), ext))
	if err != nil {
		return err
	}
	return nil
}

// Create move away the old file before create the new one.
func Create(fp string) (*os.File, error) {
	if err := MoveAway(fp); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(filepath.Dir(fp), 0600); err != nil {
		return nil, err
	}
	file, err := os.Create(fp)
	if err != nil {
		return nil, err
	}
	return file, err
}
