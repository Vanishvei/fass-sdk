package requests

// File       : snapshot.go
// Path       : requests
// Time       : CST 2023/4/26 10:57
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"github.com/Vanishvei/fass-sdk/horizontal"

	parameters "github.com/Vanishvei/fass-sdk-parameters"
	responses "github.com/Vanishvei/fass-sdk-responses"
)

func RetrieveSnapshot(parameter *parameters.RetrieveSnapshot, requestId string) (
	*responses.RetrieveSnapshot, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.RetrieveSnapshot{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func ListSnapshot(parameter *parameters.ListSnapshot, requestId string) (
	*responses.ListSnapshot, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.ListSnapshot{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func CreateSnapshot(parameter *parameters.CreateSnapshot, requestId string) (
	*responses.CreateSnapshot, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("snapshot")
	_request.SetBody(parameter)
	_request.SetMethodPOST()

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.CreateSnapshot{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RevertSnapshot(parameter *parameters.RevertSnapshot, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetBody(parameter)
	_request.SetPath("snapshot/revert")
	_request.SetMethodPUT()

	_, err = _client.callApi(_request)
	return err
}

func DeleteSnapshot(parameter *parameters.DeleteSnapshot, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodDELETE()

	_, err = _client.callApi(_request)
	return err
}
