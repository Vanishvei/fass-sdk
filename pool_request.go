package requests

// File       : pool.go
// Path       : requests
// Time       : CST 2023/4/24 15:35
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"github.com/Vanishvei/fass-sdk/horizontal"

	parameters "github.com/Vanishvei/fass-sdk-parameters"
	responses "github.com/Vanishvei/fass-sdk-responses"
)

func ListPool(parameter *parameters.ListPoolParameter, requestId string) (
	*responses.ListPoolResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetQuery(parameter.GetQuery())
	_request.SetPath("pool")

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.ListPoolResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrievePool(parameter *parameters.RetrievePoolParameter, requestId string) (
	*responses.RetrievePoolResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.RetrievePoolResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func CreatePool(parameter *parameters.CreatePoolParameter, requestId string) (
	*responses.CreatePoolResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("pool")
	_request.SetMethodPOST()
	_request.SetBody(parameter)

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.CreatePoolResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeletePool(parameter *parameters.DeletePoolParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetMethodDELETE()
	_request.SetPath(parameter.GetPath())
	_request.SetQuery(parameter.GetQuery())

	_, err = _client.CallApi(_request)
	return err
}
