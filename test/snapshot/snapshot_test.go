package snapshot

// File       : snapshot_test.go
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
	poolName            = "fast_pool"
	subsysName          = "s1000"
	volumeName          = "v1000"
	snapshotName        = "snap1000"
	invalidSnapshotName = "snap9999"
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

func TestCreateSnapshot(t *testing.T) {
	parameter := parameters.CreateSnapshot{}
	parameter.SetVolumeName(volumeName)
	parameter.SetSnapshotName(snapshotName)
	_, err := fassSDK.CreateSnapshot(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveSnapshot(t *testing.T) {
	parameter := parameters.RetrieveSnapshot{}
	parameter.SetVolumeName(volumeName)
	parameter.SetSnapshotName(snapshotName)
	_, err := fassSDK.RetrieveSnapshot(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestRetrieveSnapshotNotExists(t *testing.T) {
	parameter := parameters.RetrieveSnapshot{}
	parameter.SetVolumeName(volumeName)
	parameter.SetSnapshotName(invalidSnapshotName)
	_, err := fassSDK.RetrieveSnapshot(&parameter, uuid.New().String())
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
	t.Fail()
}

func TestListSnapshot(t *testing.T) {
	parameter := parameters.ListSnapshot{}
	parameter.SetVolumeName(volumeName)
	_, err := fassSDK.ListSnapshot(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestRevertSnapshot(t *testing.T) {
	parameter := parameters.RevertSnapshot{}
	parameter.SetVolumeName(volumeName)
	parameter.SetSnapshotName(snapshotName)
	err := fassSDK.RevertSnapshot(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestDeleteSnapshot(t *testing.T) {
	parameter := parameters.DeleteSnapshot{}
	parameter.SetVolumeName(volumeName)
	parameter.SetSnapshotName(snapshotName)
	err := fassSDK.DeleteSnapshot(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
