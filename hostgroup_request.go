package requests

// File       : hostgroup_request.go
// Path       :
// Time       : CST 2023/12/13 11:11
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"github.com/Vanishvei/fass-sdk/horizontal"

	parameters "github.com/Vanishvei/fass-sdk-parameters"
	responses "github.com/Vanishvei/fass-sdk-responses"
)

func ListHostGroup(parameters *parameters.ListHostGroup, requestId string) (
	*responses.ListHostGroup, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("acl/host_group")
	_request.SetQuery(parameters.GetQuery())

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.ListHostGroup{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteHostGroup(parameters *parameters.DeleteHostGroup, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameters.GetPath())
	_request.SetMethodDELETE()

	_, err = _client.callApi(_request)
	return err
}

func AddHostToHostGroup(parameter *parameters.AddHostToHostGroup, requestId string) (
	*responses.AddHostToHostGroup, error) {
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

	data := &responses.AddHostToHostGroup{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RemoveHostFromHostGroup(parameter *parameters.RemoveHostFromHostGroup, requestId string) (
	*responses.RemoveHostFromHostGroup, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}
	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()
	_request.SetBody(parameter)

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.RemoveHostFromHostGroup{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err

}
