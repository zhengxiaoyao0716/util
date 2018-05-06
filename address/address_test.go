package address

import (
	"fmt"
	"log"
	"testing"
)

func TestScanIPv4(t *testing.T) {
	ips, err := ScanIPv4()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ips)
}

func TestFindPorts(t *testing.T) {
	cont := make(chan bool)
	go FindPorts("localhost", 2015, func(port int, ok bool) bool {
		fmt.Println(port, ok)
		return <-cont
	})
	for i := 0; i < 4; i++ {
		fmt.Print(i, ": ")
		cont <- false
	}
	cont <- true
}

func TestFindPort(t *testing.T) {
	port := FindPort("localhost", 4000)
	fmt.Println(port)
}
