package requests

// File       : acl.go
// Path       : requests
// Time       : CST 2023/5/5 9:53
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"github.com/Vanishvei/fass-sdk/horizontal"

	parameters "github.com/Vanishvei/fass-sdk-parameters"
	responses "github.com/Vanishvei/fass-sdk-responses"
)

func ListAccount(parameter *parameters.ListAccountParameter, requestId string) (
	*responses.ListAccountResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetQuery(parameter.GetQuery())
	_request.SetPath("acl/account")

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.ListAccountResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func CreateAccount(parameter *parameters.CreateAccountParameter, requestId string) (
	*responses.CreateAccountResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("acl/account")
	_request.SetMethodPOST()
	_request.SetBody(parameter)

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.CreateAccountResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrieveAccount(parameters *parameters.RetrieveAccountParameter, requestId string) (
	*responses.RetrieveAccountResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameters.GetPath())

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.RetrieveAccountResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteAccount(parameters *parameters.DeleteAccountParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameters.GetPath())

	_, err = _client.CallApi(_request)
	return err
}

func ListGroup(parameters *parameters.ListAccountParameter, requestId string) (
	*responses.ListGroupResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("acl/group")
	_request.SetQuery(parameters.GetQuery())

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.ListGroupResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrieveGroup(parameters *parameters.RetrieveGroupParameter, requestId string) (
	*responses.RetrieveGroupResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameters.GetPath())

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.RetrieveGroupResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteGroup(parameters *parameters.DeleteGroupParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameters.GetPath())
	_request.SetMethodDELETE()

	_, err = _client.CallApi(_request)
	return err
}

func AddQualifierToGroup(parameter *parameters.AddQualifierToGroupParameter, requestId string) (
	*responses.AddQualifierToGroupResponse, error) {
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

	data := &responses.AddQualifierToGroupResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RemoveQualifierFromGroup(parameter *parameters.RemoveQualifierFromGroupParameter, requestId string) (
	*responses.RemoveQualifierFromGroupResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}
	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()
	_request.SetBody(parameter)

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.RemoveQualifierFromGroupResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err

}
