package subsys

// File       : subsys_test.go
// Path       : requests
// Time       : CST 2023/4/26 11:06
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"

	fassSDK "github.com/Vanishvei/fass-sdk"
	parameters "github.com/Vanishvei/fass-sdk-parameters"
	_ "github.com/Vanishvei/fass-sdk-responses"
	_ "github.com/Vanishvei/fass-sdk/test"
)

var (
	poolName          = "fast_pool"
	subsysName        = "s1000"
	subsysName1       = "s2000"
	subsysName2       = "s3000"
	subsysName4       = "s4000"
	invalidSubsysName = "s9999"
	volumeName        = "v1000"
	snapshotName      = "s1"
	accountName       = "account_1"
	password          = "admin@123456"
	hostGroupName     = "group_1"

	addIQN = map[string]string{
		"iqn.1994-05.com.redhat:7f11687c3ce1": "client1",
	}
	addNQN = map[string]string{
		"nqn.2014-08.org.nvmexpress:uuid:c013a7f3-e873-46a3-87d8-d5aeccd27732": "client1",
	}

	addVLANList    = []string{"172.18.0.*", "172.19.10.0/24"}
	removeVLANList = []string{"172.18.0.*"}
)

func setup() {
	createAccountParameter := parameters.CreateAccount{}
	createAccountParameter.SetAccountName(accountName)
	createAccountParameter.SetPassword(password)

	_, err := fassSDK.CreateAccount(&createAccountParameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		var ex *fassSDK.SDKError
		if errors.As(err, &ex) {
			if *(ex).Code != 800001 {
				fmt.Printf("%s", err.Error())
				panic(fmt.Sprintf("create account %s failed\n", accountName))
			}
		}
	}

	createGroupParameter := parameters.AddHostToHostGroup{}
	createGroupParameter.SetHostGroupName(hostGroupName)
	createGroupParameter.SetIQN(addIQN)
	createGroupParameter.SetNQN(addNQN)

	_, err = fassSDK.AddHostToHostGroup(&createGroupParameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		var ex *fassSDK.SDKError
		if errors.As(err, &ex) {
			if *(ex).Code != 810002 {
				fmt.Printf("%s", err.Error())
				panic(fmt.Sprintf("add client to host group %s failed\n", hostGroupName))
			}
		}
	}

	time.Sleep(2 * time.Second)
	createSubsysParameter := parameters.CreateSubsys{}
	createSubsysParameter.SetPoolName(poolName)
	createSubsysParameter.SetCapacityGB(10)
	createSubsysParameter.SetSectorSize4096()
	createSubsysParameter.SetName(subsysName4)
	_, err = fassSDK.CreateSubsys(&createSubsysParameter, uuid.New().String())
	if err != nil {
		var ex *fassSDK.SDKError
		ok := errors.As(err, &ex)
		if ok {
			if !ex.IsExists() {
				panic(fmt.Sprintf("create subsys %s failed due to %s exists\n", subsysName4, subsysName4))
			}
		}
		panic(fmt.Sprintf("create subsys %s failed due to %s\n", subsysName4, err.Error()))
	} else {
		fmt.Printf("create subsys %s success\n", subsysName4)
	}
	time.Sleep(3 * time.Second)
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

	deleteSubsysParameter := parameters.DeleteSubsys{}
	deleteSubsysParameter.SetSubsysName(subsysName4)
	deleteSubsysParameter.ForceDelete()
	deleteSubsysParameter.DeleteVolume()

	err = fassSDK.DeleteSubsys(&deleteSubsysParameter, uuid.New().String())
	if err != nil {
		fmt.Printf("delete subsys %s failed due to %s\n", subsysName4, err.Error())
	} else {
		fmt.Printf("delete subsys %s success\n", subsysName4)
	}
}

func TestCreateSubsys(t *testing.T) {
	parameter := parameters.CreateSubsys{}
	parameter.SetName(subsysName)
	parameter.SetPoolName(poolName)
	parameter.SetVolumeName(volumeName)
	parameter.SetIops(2000)
	parameter.SetBpsMB(1000)
	parameter.SetCapacityGB(10)
	parameter.SetStripeWidth4Shift128k()
	parameter.SetSectorSize4096()
	parameter.SetFormatROW()
	parameter.EnableISCSI()

	resp, err := fassSDK.CreateSubsys(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(resp.SubsysName, subsysName) {
		fmt.Printf("wrong response")
		t.FailNow()
	}
}

func TestRetrieveSubsys(t *testing.T) {
	parameter := parameters.RetrieveSubsys{}
	parameter.SetSubsysName(subsysName)

	resp, err := fassSDK.RetrieveSubsys(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(resp.SubsysName, subsysName) {
		fmt.Printf("wrong response")
		t.FailNow()
	}
}

func TestDuplicateCreateSubsys(t *testing.T) {
	parameter := parameters.CreateSubsys{}
	parameter.SetName(subsysName)
	parameter.SetPoolName(poolName)
	parameter.SetCapacityGB(1)

	_, err := fassSDK.CreateSubsys(&parameter, uuid.New().String())
	if reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	var ex *fassSDK.SDKError
	ok := errors.As(err, &ex)
	if ok {
		if ex.IsExists() {
			return
		}
		fmt.Printf("wrong response")
		t.FailNow()
	}
	fmt.Printf("%s", err.Error())
	t.FailNow()
}

func TestRetrieveSubsysNotExists(t *testing.T) {
	parameter := parameters.RetrieveSubsys{}
	parameter.SetSubsysName(invalidSubsysName)

	_, err := fassSDK.RetrieveSubsys(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		var fe *fassSDK.SDKError
		ok := errors.As(err, &fe)
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

	resp, err := fassSDK.CreateSubsysFromVolume(&parameter, requestId)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
		return
	}

	if !reflect.DeepEqual(resp.SubsysName, subsysName1) {
		fmt.Printf("wrong response")
		t.FailNow()
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

	resp, err := fassSDK.CreateSubsysFromSnapshot(&parameter, requestId)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
		return
	}

	if !reflect.DeepEqual(resp.SubsysName, subsysName2) {
		fmt.Printf("wrong response")
		t.FailNow()
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
	parameter.SetPageSize(1)

	resp, err := fassSDK.ListSubsys(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if len(resp.Token) == 0 {
		return
	}

	token, err := uuid.Parse(resp.Token)
	if err != nil {
		fmt.Printf("Invalid token %s", resp.Token)
		t.FailNow()
	}

	parameter.SetPageNum(2)
	parameter.SetPageToken(token)
	resp, err = fassSDK.ListSubsys(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if len(resp.Data) != 1 {
		fmt.Printf("The amount of data is incorrect")
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

func TestRemoveSubsysChap(t *testing.T) {
	parameter := parameters.RemoveSubsysChap{}
	parameter.SetSubsysName(subsysName)

	err := fassSDK.RemoveSubsysChap(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
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

func TestSubsysAddVLAN(t *testing.T) {
	parameter := parameters.SubsysAddVLAN{}
	parameter.SetSubsysName(subsysName)
	parameter.SetVLANList(addVLANList)

	err := fassSDK.SubsysAddVLAN(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveSubsysVLAN(t *testing.T) {
	parameter := parameters.RetrieveSubsysVLAN{}
	parameter.SetSubsysName(subsysName)

	resp, err := fassSDK.RetrieveSubsysVLAN(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(len(resp.VLANList), 2) {
		fmt.Printf("wrong vlan count")
		t.FailNow()
	}
}

func TestSubsysRemoveVLAN(t *testing.T) {
	requestId := uuid.New().String()
	parameter := parameters.SubsysRemoveVLAN{}
	parameter.SetSubsysName(subsysName)
	parameter.SetVLANList(removeVLANList)

	err := fassSDK.SubsysRemoveVLAN(&parameter, requestId)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	confirmParameter := parameters.RetrieveSubsysVLAN{}
	confirmParameter.SetSubsysName(subsysName)

	resp, err := fassSDK.RetrieveSubsysVLAN(&confirmParameter, requestId)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("confirm vlan failed due to %s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(len(resp.VLANList), 1) {
		fmt.Printf("wrong vlan count")
		t.FailNow()
	}
}

func TestDeleteSubsysVLAN(t *testing.T) {
	parameter := parameters.DeleteSubsysVLAN{}
	parameter.SetSubsysName(subsysName)

	err := fassSDK.DeleteSubsysVLAN(&parameter, uuid.New().String())
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
