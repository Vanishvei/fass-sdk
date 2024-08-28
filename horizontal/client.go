package horizontal

// File       : client.go
// Path       : horizontal
// Time       : CST 2023/4/10 15:06
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewRuntimeObject(runtime map[string]interface{}) *RuntimeObject {
	if runtime == nil {
		return &RuntimeObject{}
	}

	runtimeObject := &RuntimeObject{
		ConnectTimeout: TransInterfaceToInt(runtime["connectTimeout"]),
		ReadTimeout:    TransInterfaceToInt(runtime["readTimeout"]),
	}

	return runtimeObject
}

var hookDo = func(fn func(req *http.Request) (*http.Response, error)) func(req *http.Request) (*http.Response, error) {
	return fn
}

func DoRequest(request *Request, requestRuntime map[string]interface{}, prohibitRetries bool) (response *Response, err error) {
	runtimeObject := NewRuntimeObject(requestRuntime)

	requestURL := fmt.Sprintf(
		"http://%s:%d/api/%s/%s",
		StringValue(request.GetEndpoint()),
		IntValue(request.GetPort()),
		StringValue(request.GetApiVersion()),
		StringValue(request.GetPath()),
	)

	queryParams := request.GetQuery()
	q := url.Values{}
	for key, value := range queryParams {
		q.Add(key, StringValue(value))
	}
	queryString := q.Encode()
	if len(queryString) > 0 {
		if strings.Contains(requestURL, "?") {
			requestURL = fmt.Sprintf("%s&%s", requestURL, queryString)
		} else {
			requestURL = fmt.Sprintf("%s?%s", requestURL, queryString)
		}
	}

	httpRequest, err := http.NewRequest(StringValue(request.GetMethod()), requestURL, request.body)
	if err != nil {
		return
	}

	client := &http.Client{}
	client.Timeout = time.Duration(IntValue(runtimeObject.ReadTimeout)) * time.Second

	for key, value := range request.Headers {
		if value == nil || key == "content-length" {
			continue
		} else if key == "host" {
			httpRequest.Header["Host"] = []string{*value}
			delete(httpRequest.Header, "host")
		} else if key == "user-agent" {
			httpRequest.Header["User-Agent"] = []string{*value}
			delete(httpRequest.Header, "user-agent")
		} else {
			httpRequest.Header[key] = []string{*value}
		}
	}

	res, err := hookDo(client.Do)(httpRequest)
	if prohibitRetries {
		if err != nil {
			return response, err
		}

		response = NewResponse(res)
		for key, value := range res.Header {
			if len(value) != 0 {
				response.Headers[strings.ToLower(key)] = String(value[0])
			}
		}
		return response, err
	}

	if err != nil {
		GlobalConfig.SwitchEndpoint()
		res, err = hookDo(client.Do)(httpRequest)
		if err != nil {
			return response, err
		}
	}

	response = NewResponse(res)
	for key, value := range res.Header {
		if len(value) != 0 {
			response.Headers[strings.ToLower(key)] = String(value[0])
		}

	}

	return response, err
}
