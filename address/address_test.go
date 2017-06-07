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

func TestFindPort(t *testing.T) {
	port := FindPort("localhost", 4000)
	fmt.Println(port)
}
