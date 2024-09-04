package requests

// File       : client.go
// Path       : client
// Time       : CST 2023/4/10 14:05
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"errors"
	"github.com/Vanishvei/fass-sdk/horizontal"

	responses "github.com/Vanishvei/fass-sdk-responses"
)

func InitConfig(endpointList *[]string, port, readTimeout, connectTimeout, backoff, retryCount *int) {
	horizontal.InitConfig(endpointList, port, readTimeout, connectTimeout, backoff, retryCount)
}

type requestParams struct {
	ProhibitRetries bool
}

type RequestParams interface {
	apply(*requestParams)
}

type funcOption struct {
	f func(*requestParams)
}

func (fdo *funcOption) apply(do *requestParams) {
	fdo.f(do)
}

func newFuncOption(f func(*requestParams)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func ProhibitRetries(prohibitRetries bool) RequestParams {
	return newFuncOption(func(rp *requestParams) {
		rp.ProhibitRetries = prohibitRetries
	})
}

func defaultOptions() requestParams {
	return requestParams{
		ProhibitRetries: false,
	}
}

type defaultRequestParams struct {
	defaultOpts requestParams
}

type fassClient struct {
	Port       *int
	Endpoint   *string
	Method     *string
	Headers    map[string]*string
	RunTime    *horizontal.RuntimeObject
	ApiVersion *string
}

func newClient() (*fassClient, error) {
	client := new(fassClient)
	err := client.init(horizontal.GlobalConfig)
	return client, err
}

type SDKError = horizontal.SDKError

func (client *fassClient) init(config *horizontal.Config) (_err error) {
	if horizontal.BoolValue(horizontal.IsUnset(config)) {
		_err = horizontal.NewSDKError(map[string]interface{}{
			"code":       2,
			"message":    "'config' can not be unset",
			"data":       nil,
			"request_id": "",
		})
		return _err
	}

	client.RunTime = &horizontal.RuntimeObject{
		ConnectTimeout: config.ConnectTimeout,
		ReadTimeout:    config.ReadTimeout,
		Backoff:        config.Backoff,
		Retry:          config.RetryCount,
	}

	client.Port = config.Port
	client.Endpoint = config.CurrentEndpoint
	client.ApiVersion = config.ApiVersion
	return nil
}

func (client *fassClient) doRequest(request *horizontal.Request, prohibitRetries bool) (_result *responses.SuzakuResponse, _err error) {
	_err = horizontal.Validate(request)
	if _err != nil {
		return _result, _err
	}

	_runtime := map[string]interface{}{
		"retry":          horizontal.IntValue(horizontal.DefaultNumber(client.RunTime.Retry, horizontal.Int(2))),
		"backoff":        horizontal.IntValue(horizontal.DefaultNumber(client.RunTime.Backoff, horizontal.Int(1))),
		"readTimeout":    horizontal.IntValue(horizontal.DefaultNumber(client.RunTime.ReadTimeout, horizontal.Int(60))),
		"connectTimeout": horizontal.IntValue(horizontal.DefaultNumber(client.RunTime.ConnectTimeout, horizontal.Int(5))),
	}

	globalQueries := make(map[string]*string)
	globalHeaders := make(map[string]*string)
	request.UpdateQuery(horizontal.Merge(globalQueries, request.GetQuery()))
	request.Headers = horizontal.Merge(map[string]*string{
		"host":         request.GetEndpoint(),
		"requestId":    request.GetRequestId(),
		"user-agent":   horizontal.String("fass-sdk-golang/v1.0"),
		"accept":       horizontal.String("application/json"),
		"content-type": horizontal.String("application/json; charset=utf-8"),
	},
		globalHeaders,
		request.Headers,
	)

	_resp := &responses.SuzakuResponse{}
	for _retryTimes := 0; horizontal.BoolValue(horizontal.AllowRetry(_runtime["retry"], horizontal.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := horizontal.GetBackoffTime(_runtime["backoff"], horizontal.Int(_retryTimes))
			if horizontal.IntValue(_backoffTime) > 0 {
				horizontal.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*responses.SuzakuResponse, error) {
			response_, _err := horizontal.DoRequest(request, _runtime, prohibitRetries)
			if _err != nil {
				return _result, _err
			}

			if *response_.StatusCode == 200 {
				_res, _err := horizontal.ReadAsJSON(response_.Body)
				if _err != nil {
					return _result, _err
				}

				_, _err = horizontal.AssertAsMap(_res)
				if _err != nil {
					return _result, _err
				}

				_ = horizontal.Convert(_res, &_result)
				return _result, _err
			}

			__result := &responses.SuzakuResponse{}
			_err = horizontal.Convert(map[string]interface{}{
				"headers": response_.Headers,
				"code":    horizontal.IntValue(response_.StatusCode),
			}, __result)

			if _err != nil {
				return __result, _err
			}

			errMsg, _ := response_.ReadBody()
			__result.Message = string(errMsg)
			return __result, errors.New(__result.Message)
		}()

		if prohibitRetries {
			break
		}

		if !horizontal.BoolValue(horizontal.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

func (client *fassClient) callApi(request *horizontal.Request, opts ...RequestParams) (_result *responses.SuzakuResponse, _err error) {
	if horizontal.BoolValue(horizontal.IsUnset(request)) {
		_err = horizontal.NewSDKError(map[string]interface{}{
			"code":       300006,
			"message":    "'params' can not be unset",
			"request_id": "",
			"data":       nil,
		})
		return _result, _err
	}

	de := &defaultRequestParams{
		defaultOpts: defaultOptions(),
	}

	for _, opt := range opts {
		opt.apply(&de.defaultOpts)
	}

	request.SetPort(client.Port)
	request.SetEndpoint(client.Endpoint)
	request.SetApiVersion(client.ApiVersion)

	_resp, _err := client.doRequest(request, de.defaultOpts.ProhibitRetries)
	if _err != nil {
		sdkError := map[string]interface{}{
			"message":    _err.Error(),
			"request_id": horizontal.StringValue(request.GetRequestId()),
			"data":       nil,
		}

		if _resp != nil {
			if _resp.Code != 0 {
				sdkError["statusCode"] = _resp.Code
			}
		}

		_err = horizontal.NewSDKError(sdkError)
		return _result, _err
	}

	if _resp.Code != 0 {
		_err = horizontal.NewSDKError(map[string]interface{}{
			"code":       _resp.Code,
			"message":    _resp.Message,
			"request_id": _resp.RequestId.String(),
			"data":       _resp.Data,
		})
		return _result, _err
	}

	return _resp, _err
}
