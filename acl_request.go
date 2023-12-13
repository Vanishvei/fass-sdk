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

func ListAccount(parameter *parameters.ListAccount, requestId string) (
	*responses.ListAccount, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetQuery(parameter.GetQuery())
	_request.SetPath("acl/account")

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.ListAccount{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func CreateAccount(parameter *parameters.CreateAccount, requestId string) (
	*responses.CreateAccount, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("acl/account")
	_request.SetMethodPOST()
	_request.SetBody(parameter)

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.CreateAccount{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrieveAccount(parameters *parameters.RetrieveAccount, requestId string) (
	*responses.RetrieveAccount, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameters.GetPath())

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.RetrieveAccount{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteAccount(parameters *parameters.DeleteAccount, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameters.GetPath())

	_, err = _client.callApi(_request)
	return err
}
