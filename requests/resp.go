package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Response Http response
type Response struct {
	*http.Response
	Ok      bool
	Content []byte
	Text    string
}

// JSON Get json data.
func (r *Response) JSON() map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal(r.Content, &data); err != nil {
		return nil
	}
	return data
}

// NotOkError if used SetOkElseError(true), all HTTP status but HTTP.OK while be auto transform to this type error.
type NotOkError struct {
	Code    int
	Status  string
	content []byte
}

func (e NotOkError) Error() string {
	var data map[string]interface{}
	if err := json.Unmarshal(e.content, &data); err != nil {
		return e.Status
	}
	reason, ok := data["reas"].(string)
	if !ok {
		return e.Status + " " + string(e.content)
	}
	return e.Status + " " + reason
}

var okElseError bool

// SetOkElseError set that if auto transform HTTP status to error while it isn't HTTP.OK(statusCode == 200).
func SetOkElseError(autoTransform bool) {
	okElseError = autoTransform
}

// IsOkElseError return that is it will be auto transform to error while HTTP status isn't HTTP.OK(statusCode == 200).
func IsOkElseError() bool {
	return okElseError
}

func decorateResp(resp *http.Response, err error) (*Response, error) {
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ok := resp.StatusCode == 200
	if !ok && okElseError {
		err = NotOkError{resp.StatusCode, resp.Status, content}
	}
	return &Response{resp, ok, content, string(content)}, err
}
