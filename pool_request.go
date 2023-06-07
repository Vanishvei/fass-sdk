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

func RetrievePool(parameter *parameters.RetrievePoolParameter, requestId string) (
	*responses.RetrievePoolResponse, error) {
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

	data := &responses.RetrievePoolResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}
