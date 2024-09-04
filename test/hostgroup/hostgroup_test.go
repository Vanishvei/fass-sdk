package hostgroup

// File       : hostgroup_test.go
// Path       : test/hostgroup
// Time       : CST 2023/12/13 15:11
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"

	fassSDK "github.com/Vanishvei/fass-sdk"
	parameters "github.com/Vanishvei/fass-sdk-parameters"
	_ "github.com/Vanishvei/fass-sdk-responses"
	_ "github.com/Vanishvei/fass-sdk/test"
)

var (
	poolName      = "fast_pool"
	subsysName    = "s1000"
	volumeName    = "v1000"
	hostGroupName = "group_1"
	addIQN        = map[string]string{
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

func setup() {
	fmt.Printf("create subsys %s\n", subsysName)
	createSubsysParameter := parameters.CreateSubsys{}
	createSubsysParameter.SetPoolName(poolName)
	createSubsysParameter.SetCapacityGB(10)
	createSubsysParameter.SetSectorSize4096()
	createSubsysParameter.SetName(subsysName)
	createSubsysParameter.SetVolumeName(volumeName)
	createSubsysParameter.EnableISCSI()
	createSubsysParameter.SetFormatROW()
	_, err := fassSDK.CreateSubsys(&createSubsysParameter, uuid.New().String())
	if err != nil {
		panic(fmt.Sprintf("create subsys %s failed due to %s\n", subsysName, err.Error()))
	}

	fmt.Printf("create subsys %s success\n", subsysName)
}

func teardown() {
	fmt.Printf("delete subsys %s\n", subsysName)
	deleteSubsysParameter := parameters.DeleteSubsys{}
	deleteSubsysParameter.SetSubsysName(subsysName)
	deleteSubsysParameter.ForceDelete()
	deleteSubsysParameter.DeleteVolume()
	err := fassSDK.DeleteSubsys(&deleteSubsysParameter, uuid.New().String())
	if err != nil {
		panic(fmt.Sprintf("delete subsys %s failed due to %s\n", subsysName, err.Error()))
	}

	fmt.Printf("delete subsys %s success\n", subsysName)
}

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

func TestListHostGroup(t *testing.T) {
	parameter := parameters.ListHostGroup{}
	_, err := fassSDK.ListHostGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestSubsysBindHostGroup(t *testing.T) {
	parameter := parameters.SubsysBindHostGroup{}
	parameter.SetSubsysName(subsysName)
	parameter.SetHostGroupName(hostGroupName)

	err := fassSDK.SubsysBindHostGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestDeleteHostGroupForbidden(t *testing.T) {
	parameter := parameters.DeleteHostGroup{}
	parameter.SetHostGroupName(hostGroupName)
	err := fassSDK.DeleteHostGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		var ex fassSDK.SDKError
		if errors.As(err, &ex) {
			if *ex.Code == 810005 || *ex.Code == 810006 {
				return
			}
		}
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestSubsysUnbindHostGroup(t *testing.T) {
	parameter := parameters.SubsysUnbindHostGroup{}
	parameter.SetSubsysName(subsysName)

	err := fassSDK.SubsysUnbindHostGroup(&parameter, uuid.New().String())
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

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
