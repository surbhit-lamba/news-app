package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"news-app/internal/constants"
	"news-app/pkg/log"
)

type requestIdentifier string

const (
	NewsApi requestIdentifier = "news_api"
)

type RequestClient struct {
	Identifier requestIdentifier
	Host       string
	Scheme     string
	Authority  string
}

type RequestConfig struct {
	Method      string
	Path        string
	Body        interface{}
	QueryParams map[string]string
	HeaderKeys  []string
	Headers     map[string]string
}

func (rc *RequestClient) MakeRequest(ctx context.Context, request *RequestConfig, responseAddr interface{}) error {
	httpReq := rc.newHTTPRequest(ctx, request.Method, request.Path)
	httpReq.SetHost(rc.Authority)
	// add body for request
	reqBytes, err := json.Marshal(request.Body)
	if err != nil {
		log.ErrorfWithContext(ctx, "%s client - request marshal error | ", rc.Identifier, err.Error())
		return fmt.Errorf("%s client - request marshal error", rc.Identifier)
	}
	httpReq.SetBody(bytes.NewBuffer(reqBytes))

	// for manual setting of headers
	for key, name := range request.Headers {
		httpReq.SetHeader(key, name)
	}

	// add all query params
	for param, value := range request.QueryParams {
		httpReq.SetQueryParam(param, value)
	}

	httpClient := NewClient(ctx)
	httpResp, err := httpClient.Execute(ctx, *httpReq)
	if err != nil {
		fields := make(map[string]interface{})
		fields[constants.ResponseAdd] = string(reqBytes)
		fields[constants.Path] = request.Path
		fields[constants.Err] = err.Error()
		log.ErrorfWithContext(ctx, "%s http request path error", err)
		return fmt.Errorf("%s http request error %+v", rc.Identifier, fields)
	}

	if httpResp.StatusCode != http.StatusOK {
		fields := make(map[string]interface{})
		fields[constants.ResponseAdd] = fmt.Sprintf("%+v", httpResp.GetBody())
		fields[constants.Path] = request.Path
		return fmt.Errorf("%s http non 200 responseAddr, status code: %v, error: %+v", rc.Identifier, httpResp.StatusCode, err)
	}
	errRes := httpResp.Bind(&responseAddr)
	if errRes != nil {
		fields := make(map[string]interface{})
		fields[constants.ResponseAdd] = fmt.Sprintf("%+v", httpResp.GetBody())
		fields[constants.Request] = string(reqBytes)
		fields[constants.Err] = errRes.Error()

		return fmt.Errorf("%s could not bind responseAddr, error: %+v", rc.Identifier, fields)
	}

	return nil
}

func (os *RequestClient) newHTTPRequest(ctx context.Context, method, path string) *Request {
	URL := url.URL{
		Scheme: os.Scheme,
		Host:   os.Host,
		Path:   path,
	}

	rq := NewRequest(ctx, method, URL)
	rq.SetHeader("Content-Type", "application/json")

	return rq
}
