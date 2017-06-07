// Package bat build and run bat commands.
// Notice that bat run async.
package bat

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Bat object
type Bat struct {
	FileName string
	*os.File
}

// Create a bat object.
func Create(name string) (*Bat, error) {
	var fileName string
	if strings.HasSuffix(name, ".bat") {
		fileName = name
	} else {
		fileName = name + ".bat"
	}
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	if _, err := file.WriteString(":: github.com/zhengxiaoyao0716/util/bat\n:: " + name + "\n"); err != nil {
		file.Close()
		return nil, err
	}
	return &Bat{fileName, file}, nil
}

// Append a line of command.
func (bat *Bat) Append(cmd string) (*Bat, error) {
	if _, err := bat.File.WriteString(cmd + "\n"); err != nil {
		return nil, err
	}
	return bat, nil
}

// Run the bat.
func (bat *Bat) Run(params ...string) error {
	paramStr := ""
	for _, param := range params {
		paramStr += param
	}
	bat.Close()
	return exec.Command(bat.FileName, paramStr).Run()
}

// Remove and free the bat.
func (bat *Bat) Remove() error {
	bat.Close()
	return os.Remove(bat.FileName)
}

// Exec execute bat command, equals to `Create && Append && Close && Run && Remove`
func Exec(cmds ...string) error {
	bat, err := Create(os.TempDir() + "\\util_bat_" + fmt.Sprint(time.Now().Unix()))
	if err != nil {
		return err
	}
	defer bat.Remove()

	for _, cmd := range cmds {
		if bat, err = bat.Append(cmd); err != nil {
			return err
		}
	}

	if err = bat.Close(); err != nil {
		return err
	}

	return exec.Command(bat.FileName).Run()
}
