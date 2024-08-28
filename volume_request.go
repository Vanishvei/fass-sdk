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

func ListVolume(parameter *parameters.ListVolume, requestId string) (
	*responses.ListVolume, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("volume")
	_request.SetQuery(parameter.GetQuery())

	resp, err := _client.callApi(_request, ProhibitRetries(true))
	if err != nil {
		return nil, err
	}

	lv := &responses.ListVolume{}
	if resp.Token != nil {
		lv.Token = resp.Token.String()
		lv.Total = *resp.Total
		lv.PageNum = *resp.PageNum
		lv.PageSize = *resp.PageSize
	}
	err = horizontal.ConvertToSuzakuResp(resp.Data, &lv.Data)
	return lv, err
}

func RetrieveVolume(parameter *parameters.RetrieveVolume, requestId string) (
	*responses.RetrieveVolume, error) {
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

	data := &responses.RetrieveVolume{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteVolume(parameter *parameters.DeleteVolume, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetQuery(parameter.GetQuery())
	_request.SetMethodDELETE()

	_, err = _client.callApi(_request)
	return err
}

func ExpandVolume(parameter *parameters.ExpandVolume, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)
	_request.SetMethodPUT()

	_, err = _client.callApi(_request)
	return err
}

func FlattenVolume(parameter *parameters.FlattenVolume, requestId string) (
	*responses.FlattenVolume, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)
	_request.SetMethodPUT()

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.FlattenVolume{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func SetQosOfVolume(parameter *parameters.SetQosOfVolume, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}
	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)
	_request.SetMethodPUT()

	_, err = _client.callApi(_request)
	return err
}

func FlattenVolumeProgress(parameter *parameters.GetFlattenVolumeProgress, requestId string) (
	*responses.FlattenVolumeProgress, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.FlattenVolumeProgress{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func StopFlattenVolume(parameter *parameters.StopFlattenVolume, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()

	_, err = _client.callApi(_request)
	return err
}
