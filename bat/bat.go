package bat

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// Bat 批处理对象
type Bat struct {
	FileName string
	*os.File
}

func init() {
	if err := os.MkdirAll("tmp/", os.ModePerm); err != nil {
		log.Fatalln("创建临时文件夹失败，" + err.Error())
	}
}

// Create 创建一个批处理
func Create(name string) (*Bat, error) {
	fileName := "tmp\\" + name + ".bat"
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

// Append 增加一行命令
func (bat *Bat) Append(cmd string) (*Bat, error) {
	if _, err := bat.File.WriteString(cmd + "\n"); err != nil {
		return nil, err
	}
	return bat, nil
}

// Run 执行bat
func (bat *Bat) Run(params ...string) error {
	paramStr := ""
	for _, param := range params {
		paramStr += param
	}
	return exec.Command(bat.FileName, paramStr).Run()
}

// Remove 销毁bat文件
func (bat *Bat) Remove() error {
	return os.Remove(bat.FileName)
}

// Exec 执行bat命令，相当于Create&Append&Close&Run&Remove
func Exec(cmds ...string) error {
	bat, err := Create("util_bat_exec_" + fmt.Sprint(time.Now().Unix()))
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
