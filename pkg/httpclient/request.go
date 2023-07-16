package httpclient

import (
	"bytes"
	"context"
	"io"
	"net/url"
)

// Request represents an individual HTTP request
type Request struct {
	context     context.Context
	method      string
	url         url.URL
	headers     map[string][]string
	queryParams map[string][]string
	host        string
	body        io.Reader
}

// NewRequest creates a new http request instance
func NewRequest(ctx context.Context, method string, url url.URL) *Request {
	queryParams := make(map[string][]string)
	headers := make(map[string][]string)

	rq := &Request{
		context:     ctx,
		method:      method,
		url:         url,
		headers:     headers,
		queryParams: queryParams,
	}

	return rq
}

// SetHost sets the host for http request
func (req *Request) SetHost(host string) {
	req.host = host
}

// SetHeader adds header for http request
func (req *Request) SetHeader(key string, value string) {
	req.headers[key] = append(make([]string, 0), value)
}

// SetQueryParam sets the given query param and value
func (req *Request) SetQueryParam(key string, value string) {
	req.queryParams[key] = append(make([]string, 0), value)

	data := url.Values(req.queryParams)
	req.url.RawQuery = data.Encode()
}

// SetBody sets body for http call
func (req *Request) SetBody(body *bytes.Buffer) {
	req.body = body
}
