package main

import (
	"github.com/zhengxiaoyao0716/util/console"
)

func main() {
	console.Log(console.ReadPass("Password: "))

	console.PushLine("Hello World.")
	console.PushLine("trigger")
	console.PushLine("abort")
	for {
		line := console.ReadLine()
		switch line {
		case "trigger":
			console.TriggerInterrupt()
		case "abort":
			console.AbortInterrupt()
		default:
			console.Log(line)
		}
	}
}
func init() {
	console.CatchInterrupt()
}
