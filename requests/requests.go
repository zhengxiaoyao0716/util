// Package requests is a simple http client, witch provide
// practical 'text()', 'json()' utils like python side-package.
package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type request func(url string, data map[string]interface{}) (*Response, error)

// Default client with HTTP request methods
var Default = New(http.DefaultClient)

// HTTP request methods
var (
	Get  = Default.Get
	Post = Default.Post
)

// New .
func New(client *http.Client) struct {
	Get  request
	Post request
} {
	return struct {
		Get  request
		Post request
	}{
		Get: func(url string, data map[string]interface{}) (*Response, error) {
			return get(client, url, data)
		},
		Post: func(url string, data map[string]interface{}) (*Response, error) {
			return post(client, url, data)
		},
	}
}

func get(client *http.Client, url string, data map[string]interface{}) (*Response, error) {
	var params string
	if data != nil {
		for key, value := range data {
			switch value := value.(type) {
			case []string:
				for _, v := range value {
					params += key + "=" + v + "&"
				}
			case string:
				params += key + "=" + value + "&"
			case []interface{}:
				for _, v := range value {
					params += key + "=" + fmt.Sprint(v) + "&"
				}
			default:
				params += key + "=" + fmt.Sprint(value) + "&"
			}
		}
	}
	return decorateResp(client.Get(url + "?" + params))
}

func post(client *http.Client, url string, data map[string]interface{}) (*Response, error) {
	var content []byte
	var err error
	var resp *http.Response
	if data != nil {
		content, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
		resp, err = client.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(content))
	} else {
		resp, err = client.Post(url, "", nil)
	}
	return decorateResp(resp, err)
}
