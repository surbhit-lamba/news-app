package httpclient

import (
	"encoding/json"
	"net/http"
)

// Response represents the response recevied from an HTTP request
type Response struct {
	body []byte
	http.Response
}

// Bind unmarshals response body to given interface
func (resp *Response) Bind(v interface{}) error {
	err := json.Unmarshal(resp.body, &v)
	if err != nil {
		return err
	}

	return nil
}

func (resp *Response) GetBody() string {
	return string(resp.body)
}
