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

func CreateSubsys(parameter *parameters.CreateSubsys, requestId string) (
	*responses.CreateSubsys, error) {
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

	data := &responses.CreateSubsys{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func CreateSubsysFromVolume(parameter *parameters.CreateSubsysFromVolume,
	requestId string) (*responses.CreateSubsys, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	// create temp snapshot
	_request := horizontal.NewRequest(requestId)
	createSnapshot := &parameters.CreateSnapshot{}
	createSnapshot.SetVolumeName(parameter.GetSourceVolumeName())
	createSnapshot.SetSnapshotName(parameter.GetSubsysName())
	_, err = CreateSnapshot(createSnapshot, requestId)
	if err != nil {
		return nil, err
	}

	// use tmp snapshot create volume
	_request.SetBody(parameter)
	_request.SetPath("subsys")
	_request.SetMethodPOST()
	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	// delete tmp snapshot
	deleteSnapshot := &parameters.DeleteSnapshot{}
	deleteSnapshot.SetVolumeName(parameter.GetSourceVolumeName())
	deleteSnapshot.SetSnapshotName(parameter.GetSubsysName())
	_ = DeleteSnapshot(deleteSnapshot, requestId)

	data := &responses.CreateSubsys{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func CreateSubsysFromSnapshot(parameter *parameters.CreateSubsysFromSnapshot,
	requestId string) (*responses.CreateSubsys, error) {
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

	data := &responses.CreateSubsys{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func ListSubsys(parameter *parameters.ListSubsys, requestId string) (
	*responses.ListSubsys, error) {
	_client, err := newClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("subsys")
	_request.SetQuery(parameter.GetQuery())

	resp, err := _client.callApi(_request)
	if err != nil {
		return nil, err
	}

	data := &responses.ListSubsys{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrieveSubsys(parameter *parameters.RetrieveSubsys, requestId string) (
	*responses.RetrieveSubsys, error) {
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

	data := &responses.RetrieveSubsys{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteSubsys(parameter *parameters.DeleteSubsys, requestId string) error {
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

func ExportSubsys(parameter *parameters.ExportSubsys, requestId string) error {
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

func UnexportSubsys(parameter *parameters.UnexportSubsys, requestId string) error {
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

func RetrieveSubsysChap(parameter *parameters.RetrieveSubsysChap, requestId string) (
	*responses.RetrieveSubsysChap, error) {
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

	data := &responses.RetrieveSubsysChap{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func SetSubsysChap(parameter *parameters.SetSubsysChap, requestId string) error {
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

func RemoveSubsysChap(parameter *parameters.RemoveSubsysChap, requestId string) error {
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

func SubsysAddVLAN(parameter *parameters.SubsysAddVLAN, requestId string) error {
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

func RetrieveSubsysVLAN(parameter *parameters.RetrieveSubsysVLAN, requestId string) (
	*responses.RetrieveSubsysVLAN, error) {
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

	data := &responses.RetrieveSubsysVLAN{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func SubsysRemoveVLAN(parameter *parameters.SubsysRemoveVLAN, requestId string) error {
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

func DeleteSubsysVLAN(parameter *parameters.DeleteSubsysVLAN, requestId string) error {
	_client, err := newClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())

	_, err = _client.callApi(_request)
	return err
}

func SubsysBindHostGroup(parameter *parameters.SubsysBindHostGroup, requestId string) error {
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

func SubsysUnbindHostGroup(parameter *parameters.SubsysUnbindHostGroup, requestId string) error {
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

func RetrieveSubsysHostGroup(parameter *parameters.RetrieveSubsysHostGroup,
	requestId string) (*responses.RetrieveSubsysHostGroup, error) {
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

	data := &responses.RetrieveSubsysHostGroup{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}
