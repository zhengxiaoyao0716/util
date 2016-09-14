package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Get HTTP Get method.
func Get(url string, data map[string]string) (*Response, error) {
	var params string
	if data != nil {
		for key, value := range data {
			params += key + "=" + value + "&"
		}
	}
	return decorateResp(http.Get(url + "?" + params))
}

// Post HTTP Post method.
func Post(url string, data map[string]interface{}) (*Response, error) {
	var content []byte
	var err error
	var resp *http.Response
	if data != nil {
		content, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
		resp, err = http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(content))
	} else {
		resp, err = http.Post(url, "", nil)
	}
	return decorateResp(resp, err)
}
