package subsys

// File       : subsys_test.go
// Path       : requests
// Time       : CST 2023/4/26 11:06
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
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
	poolName          = "pool1"
	subsysName        = "s1000"
	subsysName1       = "s2000"
	subsysName2       = "s3000"
	invalidSubsysName = "s9999"
	volumeName        = "v1000"
	snapshotName      = "s1"
	accountName       = "account_1"
	password          = "admin@1234"
	hostGroupName     = "group_1"

	addIQN = map[string]string{
		"iqn.1994-05.com.redhat:7f11687c3ce1": "client1",
	}
	addNQN = map[string]string{
		"nqn.2014-08.org.nvmexpress:uuid:c013a7f3-e873-46a3-87d8-d5aeccd27732": "client1",
	}
)

func setup() {
	createAccountParameter := parameters.CreateAccount{}
	createAccountParameter.SetAccountName(accountName)
	createAccountParameter.SetPassword(password)

	_, err := fassSDK.CreateAccount(&createAccountParameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		panic(fmt.Sprintf("create account %s failed\n", accountName))
	}

	createGroupParameter := parameters.AddHostToHostGroup{}
	createGroupParameter.SetHostGroupName(hostGroupName)
	createGroupParameter.SetIQN(addIQN)
	createGroupParameter.SetNQN(addNQN)

	_, err = fassSDK.AddHostToHostGroup(&createGroupParameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		panic(fmt.Sprintf("add client to host group %s failed\n", hostGroupName))
	}
}

func teardown() {
	deleteGroupParameter := parameters.DeleteHostGroup{}
	deleteGroupParameter.SetHostGroupName(hostGroupName)
	err := fassSDK.DeleteHostGroup(&deleteGroupParameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("delete host group %s failed\n", hostGroupName)
	}

	deleteAccountParameter := parameters.DeleteAccount{}
	deleteAccountParameter.SetAccountName(accountName)

	err = fassSDK.DeleteAccount(&deleteAccountParameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("delete account %s failed\n", accountName)
	}
}

func TestCreateSubsys(t *testing.T) {
	parameter := parameters.CreateSubsys{}
	parameter.SetName(subsysName)
	parameter.SetPoolName(poolName)
	parameter.SetVolumeName(volumeName)
	parameter.SetSectorSize4096()
	parameter.SetCapacityGB(10)
	parameter.EnableISCSI()
	parameter.SetBpsMB(1000)
	parameter.SetIops(2000)
	parameter.SetFormatROW()

	_, err := fassSDK.CreateSubsys(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveSubsys(t *testing.T) {
	parameter := parameters.RetrieveSubsys{}
	parameter.SetSubsysName(subsysName)

	_, err := fassSDK.RetrieveSubsys(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveSubsysNotExists(t *testing.T) {
	parameter := parameters.RetrieveSubsys{}
	parameter.SetSubsysName(invalidSubsysName)

	_, err := fassSDK.RetrieveSubsys(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fe, ok := err.(*fassSDK.SDKError)
		if ok {
			if !fe.IsNotExists() {
				t.Fail()
				return
			}
			if fe.IsExists() {
				t.Fail()
				return
			}
			return
		}
	}
	fmt.Printf("%s", err.Error())
	t.FailNow()
}

func TestCreateSubsysFromVolume(t *testing.T) {
	requestId := uuid.New().String()

	parameter := parameters.CreateSubsysFromVolume{}
	parameter.SetName(subsysName1)
	parameter.SetPoolName(poolName)
	parameter.SetSrcVolumeName(volumeName)
	parameter.InheritQos()
	parameter.EnableISCSI()

	_, err := fassSDK.CreateSubsysFromVolume(&parameter, requestId)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
		return
	}

	err = deleteSubsys(subsysName1)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestCreateSubsysFromSnapshot(t *testing.T) {
	requestId := uuid.New().String()
	createSnapshotParameter := parameters.CreateSnapshot{}
	createSnapshotParameter.SetVolumeName(volumeName)
	createSnapshotParameter.SetSnapshotName(snapshotName)
	_, err := fassSDK.CreateSnapshot(&createSnapshotParameter, requestId)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
		return
	}

	parameter := parameters.CreateSubsysFromSnapshot{}
	parameter.SetName(subsysName2)
	parameter.SetPoolName(poolName)
	parameter.SetSrcVolumeName(volumeName)
	parameter.SetSrcSnapshotName(snapshotName)
	parameter.InheritQos()
	parameter.EnableISCSI()

	_, err = fassSDK.CreateSubsysFromSnapshot(&parameter, requestId)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
		return
	}

	deleteSnapshotParameter := parameters.DeleteSnapshot{}
	deleteSnapshotParameter.SetVolumeName(volumeName)
	deleteSnapshotParameter.SetSnapshotName(snapshotName)

	err = fassSDK.DeleteSnapshot(&deleteSnapshotParameter, requestId)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	err = deleteSubsys(subsysName2)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestListSubsys(t *testing.T) {
	parameter := parameters.ListSubsys{}

	_, err := fassSDK.ListSubsys(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestExportSubsys(t *testing.T) {
	parameter := parameters.ExportSubsys{}
	parameter.SetSubsysName(subsysName)
	parameter.ExportISCSI()

	err := fassSDK.ExportSubsys(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestUnexportSubsys(t *testing.T) {
	parameter := parameters.UnexportSubsys{}
	parameter.SetSubsysName(subsysName)
	parameter.UnexportISCSI()

	err := fassSDK.UnexportSubsys(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestSetSubsysAuth(t *testing.T) {
	parameter := parameters.SubsysBindHostGroup{}
	parameter.SetSubsysName(subsysName)
	parameter.SetHostGroupName(hostGroupName)

	err := fassSDK.SubsysBindHostGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveSubsysHostGroup(t *testing.T) {
	parameter := parameters.RetrieveSubsysHostGroup{}
	parameter.SetSubsysName(subsysName)

	resp, err := fassSDK.RetrieveSubsysHostGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(resp.HostGroup, hostGroupName) {
		fmt.Printf("auth result %s is not equal to %s\n", resp.HostGroup, hostGroupName)
		t.FailNow()
	}
}

func TestSetSubsysChap(t *testing.T) {
	parameter := parameters.SetSubsysChap{}
	parameter.SetSubsysName(subsysName)
	parameter.SetAccountName(accountName)

	err := fassSDK.SetSubsysChap(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveSubsysChap(t *testing.T) {
	parameter := parameters.RetrieveSubsysChap{}
	parameter.SetSubsysName(subsysName)

	resp, err := fassSDK.RetrieveSubsysChap(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(resp.Chap, accountName) {
		fmt.Printf("chap result %s is not equal to %s\n", resp.Chap, accountName)
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

func TestRemoveSubsysChap(t *testing.T) {
	parameter := parameters.RemoveSubsysChap{}
	parameter.SetSubsysName(subsysName)

	err := fassSDK.RemoveSubsysChap(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestDeleteSubsys(t *testing.T) {
	err := deleteSubsys(subsysName)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func deleteSubsys(subsysName_ string) error {
	parameter := parameters.DeleteSubsys{}
	parameter.SetSubsysName(subsysName_)
	parameter.DeleteVolume()
	parameter.ForceDelete()

	return fassSDK.DeleteSubsys(&parameter, uuid.New().String())
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
