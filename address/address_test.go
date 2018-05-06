package address

import (
	"fmt"
	"log"
	"testing"
)

func TestScanNets(t *testing.T) {
	netMap, err := ScanNets()
	if err != nil {
		log.Fatalln(err)
	}
	for name, nets := range netMap {
		fmt.Printf("[%s]\n", name)
		for _, net := range nets {
			fmt.Println(net)
		}
		fmt.Println()
	}
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
