package requests

// File       : volume.go
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

func ListVolume(parameter *parameters.ListVolumeParameter, requestId string) (
	*responses.ListVolumeResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("volume")
	_request.SetQuery(parameter.GetQuery())

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.ListVolumeResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrieveVolume(parameter *parameters.RetrieveVolumeParameter, requestId string) (
	*responses.RetrieveVolumeResponse, error) {
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

	data := &responses.RetrieveVolumeResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteVolume(parameter *parameters.DeleteVolumeParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetQuery(parameter.GetQuery())
	_request.SetMethodDELETE()

	_, err = _client.CallApi(_request)
	return err
}

func ExpandVolume(parameter *parameters.ExpandVolumeParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)
	_request.SetMethodPUT()

	_, err = _client.CallApi(_request)
	return err
}

func FlattenVolume(parameter *parameters.FlattenVolumeParameter, requestId string) (
	*responses.FlattenVolumeResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)
	_request.SetMethodPUT()

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.FlattenVolumeResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func SetQosOfVolume(parameter *parameters.SetQosOfVolumeParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}
	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)
	_request.SetMethodPUT()

	_, err = _client.CallApi(_request)
	return err
}

func FlattenVolumeProgress(parameter *parameters.GetFlattenVolumeProgress, requestId string) (
	*responses.FlattenVolumeProgressResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.FlattenVolumeProgressResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func StopFlattenVolume(parameter *parameters.StopFlattenVolumeParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())

	_, err = _client.CallApi(_request)
	return err
}
