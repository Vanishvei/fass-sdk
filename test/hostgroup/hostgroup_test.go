package hostgroup

// File       : hostgroup_test.go
// Path       : test/hostgroup
// Time       : CST 2023/12/13 15:11
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"

	fassSDK "github.com/Vanishvei/fass-sdk"
	parameters "github.com/Vanishvei/fass-sdk-parameters"
)

var (
	hostGroupName = "group_1"

	addIQN = map[string]string{
		"iqn.1994-05.com.redhat:7f11687c3ce1": "client1",
		"iqn.1994-05.com.redhat:7f21671c3ce2": "client2",
	}
	addNQN = map[string]string{
		"nqn.2014-08.org.nvmexpress:uuid:c013a7f3-e873-46a3-87d8-d5aeccd27732": "client1",
		"nqn.2014-08.org.nvmexpress:uuid:c013a7f3-e873-46a3-87d8-d5aeccd12345": "client2",
	}

	removeIQN = map[string]string{
		"iqn.1994-05.com.redhat:7f11687c3ce1": "client1",
	}
	removeNQN = map[string]string{
		"nqn.2014-08.org.nvmexpress:uuid:c013a7f3-e873-46a3-87d8-d5aeccd27732": "client1",
	}
)

func TestAddHostToHostGroup(t *testing.T) {
	parameter := parameters.AddHostToHostGroup{}
	parameter.SetIQN(addIQN)
	parameter.SetNQN(addNQN)
	parameter.SetHostGroupName(hostGroupName)
	_, err := fassSDK.AddHostToHostGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestListSubsysOfHostGroup(t *testing.T) {
	parameter := parameters.ListSubsysOfHostGroup{}
	parameter.SetHostGroupName(hostGroupName)

	_, err := fassSDK.ListSubsysOfHostGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestListHostGroup(t *testing.T) {
	parameter := parameters.ListHostGroup{}
	_, err := fassSDK.ListHostGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRemoveHostFromHostGroup(t *testing.T) {
	parameter := parameters.RemoveHostFromHostGroup{}
	parameter.SetIQN(removeIQN)
	parameter.SetNQN(removeNQN)
	parameter.SetHostGroupName(hostGroupName)
	_, err := fassSDK.RemoveHostFromHostGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestDeleteHostGroup(t *testing.T) {
	parameter := parameters.DeleteHostGroup{}
	parameter.SetHostGroupName(hostGroupName)
	err := fassSDK.DeleteHostGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}
