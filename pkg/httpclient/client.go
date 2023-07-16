package httpclient

import (
	"compress/gzip"
	"context"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"news-app/internal/constants"
	"time"
)

const HTTP_STATUS_OK = 200

// Client wraps and http client
type Client struct {
	Client *http.Client
}

type BatchExecutionJob struct {
	BatchIndex int
	Request    Request
}

type BatchError struct {
	BatchIndex int
	Error      error
}

// NewClient return a simple http client
func NewClient(ctx context.Context) *Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DialContext = (&net.Dialer{
		Timeout:   time.Duration(5) * time.Second,
		KeepAlive: time.Duration(20) * time.Second,
	}).DialContext
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 50
	t.MaxIdleConnsPerHost = 50

	client := Client{
		Client: &http.Client{
			Timeout:   5 * time.Minute,
			Transport: t,
		},
	}

	return &client
}

// Execute executes the given request object and returns a repsonse object
func (c *Client) Execute(ctx context.Context, req Request) (Response, error) {
	hreq, err := http.NewRequestWithContext(ctx, req.method, req.url.String(), req.body)
	if err != nil {
		// TODO : Do logging/New Relic
		return Response{}, err
	}

	hreq.Header = req.headers
	if req.host != "" {
		hreq.Host = req.host
	}

	// make the request
	resp, err := c.Client.Do(hreq)
	if err != nil {
		// TODO : Do logging/New Relic
		return Response{}, err
	}
	defer resp.Body.Close()
	var respBody []byte
	if resp.Header.Get(constants.HTTPHeaderContentEncoding) == constants.ContentEncodingGZIP {
		reader, parsingError := gzip.NewReader(resp.Body)
		if parsingError != nil {
			return Response{}, parsingError
		}
		defer reader.Close()
		respBody, parsingError = io.ReadAll(reader)
		if parsingError != nil {
			return Response{}, parsingError
		}
	} else {
		respBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return Response{}, err
		}
	}

	if resp.StatusCode != HTTP_STATUS_OK {
		// TODO : Do logging/New Relic
	}

	response := Response{
		body:     respBody,
		Response: *resp,
	}

	return response, nil
}
