package requests

import (
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := Get("http://localhost:5000/api/s/user/list", nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Content)
	log.Println(resp.Text)
	log.Println(resp.JSON())
}

func TestPost(t *testing.T) {
	resp, err := Post("http://localhost:5000/api/s/user/join", map[string]interface{}{"address": "http://localhost:5000"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Content)
	log.Println(resp.Text)
	log.Println(resp.JSON())
}

func TestError(t *testing.T) {
	SetOkElseError(true)
	resp, err := Post("http://localhost:5000/api/a/connect/create", map[string]interface{}{"address": "test error."})
	if err != nil {
		log.Println(err)
	}
	log.Println(resp.StatusCode)
	log.Println(resp.Content)
	log.Println(resp.Text)
	log.Println(resp.JSON())
}
func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
