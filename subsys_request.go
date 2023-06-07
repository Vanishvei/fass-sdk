package requests

// File       : subsys.go
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

func CreateSubsys(parameter *parameters.CreateSubsysParameter, requestId string) (
	*responses.CreateSubsysResponse, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetBody(parameter)
	_request.SetPath("subsys")
	_request.SetMethodPOST()
	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.CreateSubsysResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func ListSubsys(parameter *parameters.ListSubsysParameter, requestId string) (
	*responses.ListSubsysResponse, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("pool")
	_request.SetQuery(parameter.GetQuery())

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.ListSubsysResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrieveSubsys(parameter *parameters.RetrieveSubsysParameter, requestId string) (
	*responses.RetrieveSubsysResponse, error) {
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

	data := &responses.RetrieveSubsysResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteSubsys(parameter *parameters.DeleteSubsysParameter, requestId string) error {
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

func ExportSubsys(parameter *parameters.ExportSubsysParameter, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()
	_request.SetBody(parameter)

	_, err = _client.callApi(_request)
	return err
}

func UnexportSubsys(parameter *parameters.UnexportSubsysParameter, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()
	_request.SetBody(parameter)

	_, err = _client.callApi(_request)
	return err
}

func RetrieveSubsysAuth(parameter *parameters.RetrieveSubsysAuthParameter,
	requestId string) (*responses.RetrieveSubsysAuthResponse, error) {
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

	data := &responses.RetrieveSubsysAuthResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func SetSubsysAuth(parameter *parameters.SetSubsysAuthParameter, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()
	_request.SetBody(parameter)

	_, err = _client.callApi(_request)
	return err
}

func RemoveSubsysAuth(parameter *parameters.RemoveSubsysAuthParameter, requestId string) error {
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

func RetrieveSubsysChap(parameter *parameters.RetrieveSubsysChapParameter, requestId string) (
	*responses.RetrieveSubsysChapResponse, error) {
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

	data := &responses.RetrieveSubsysChapResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func SetSubsysChap(parameter *parameters.SetSubsysChapParameter, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()
	_request.SetBody(parameter)

	_, err = _client.callApi(_request)
	return err
}

func RemoveSubsysChap(parameter *parameters.RemoveSubsysChapParameter, requestId string) error {
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
