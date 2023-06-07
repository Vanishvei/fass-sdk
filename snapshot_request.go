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

func RetrieveSnapshot(parameter *parameters.RetrieveSnapshotParameter, requestId string) (
	*responses.RetrieveSnapshotResponse, error) {
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

	data := &responses.RetrieveSnapshotResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func ListSnapshot(parameter *parameters.ListSnapshotParameter, requestId string) (
	*responses.ListSnapshotResponse, error) {
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

	data := &responses.ListSnapshotResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func CreateSnapshot(parameter *parameters.CreateSnapshotParameter, requestId string) (
	*responses.CreateSnapshotResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("snapshot")
	_request.SetBody(parameter)
	_request.SetMethodPOST()

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.CreateSnapshotResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RevertSnapshot(parameter *parameters.RevertSnapshotParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetBody(parameter)
	_request.SetPath("snapshot")
	_request.SetMethodPUT()

	_, err = _client.CallApi(_request)
	return err
}

func DeleteSnapshot(parameter *parameters.DeleteSnapshotParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodDELETE()

	_, err = _client.CallApi(_request)
	return err
}
