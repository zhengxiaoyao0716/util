package requests

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

var address = "localhost:4000"
var api = "http://" + address + "/api/ping"

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	go func() {
		http.HandleFunc("/api/ping", func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte(fmt.Sprintf(`{"value": "%s"}`, time.Now().Local())))
		})
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Fatalln(err)
		}
	}()

	err := errors.New("wait for server.")
	for err != nil {
		_, err = http.Head(api)
		fmt.Print(".")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\n> Test server started.")
}

func TestGet(t *testing.T) {
	resp, err := Get(api, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(resp.Content)
	// fmt.Println(resp.Text)
	fmt.Println(resp.JSON())
}

func TestPost(t *testing.T) {
	resp, err := Post(api, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(resp.Content)
	// fmt.Println(resp.Text)
	fmt.Println(resp.JSON())
}

func TestError(t *testing.T) {
	SetOkElseError(true)
	resp, err := Post(api+"/notfound", nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Content)
	// fmt.Println(resp.Text)
	fmt.Println(resp.JSON())
}
