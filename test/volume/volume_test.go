package volume

// File       : volume_test.go
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
	"time"

	"github.com/google/uuid"

	fassSDK "github.com/Vanishvei/fass-sdk"
	parameters "github.com/Vanishvei/fass-sdk-parameters"
	responses "github.com/Vanishvei/fass-sdk-responses"
	_ "github.com/Vanishvei/fass-sdk/test"
)

var (
	poolName          = "pool1"
	srcSubsysName     = "s1000"
	srcVolumeName     = "v1000"
	srcSnapshotName   = "snap1000"
	newSubsysName1    = "s2000"
	newVolumeName1    = "v2000"
	newSubsysName2    = "s3000"
	newVolumeName2    = "v3000"
	invalidVolumeName = "v9999"
	taskId            = ""
)

func setup() {
	fmt.Printf("create source subsys %s\n", srcSubsysName)
	createSubsysParameter := parameters.CreateSubsys{}
	createSubsysParameter.SetPoolName(poolName)
	createSubsysParameter.SetCapacityGB(10)
	createSubsysParameter.SetSectorSize4096()
	createSubsysParameter.SetName(srcSubsysName)
	createSubsysParameter.SetVolumeName(srcVolumeName)
	createSubsysParameter.EnableISCSI()
	createSubsysParameter.SetFormatROW()
	_, err := fassSDK.CreateSubsys(&createSubsysParameter, uuid.New().String())
	if err != nil {
		panic(fmt.Sprintf("create source subsys %s failed due to %s\n", srcSubsysName, err.Error()))
	}
	fmt.Printf("create source subsys %s success\n", srcSubsysName)

	time.Sleep(3 * time.Second)

	fmt.Printf("create source snapshot %s\n", srcSnapshotName)
	createSnapshotParameter := parameters.CreateSnapshot{}
	createSnapshotParameter.SetVolumeName(srcVolumeName)
	createSnapshotParameter.SetSnapshotName(srcSnapshotName)

	_, err = fassSDK.CreateSnapshot(&createSnapshotParameter, uuid.New().String())
	if err != nil {
		panic(fmt.Sprintf("create source snapshot %s failed due to %s\n", srcSnapshotName, err.Error()))
	}

	fmt.Printf("create source snapshot %s success\n", srcSnapshotName)

	err = createVolumeFromSnapshot(newSubsysName2, newVolumeName2)
	if err != nil {
		panic(fmt.Sprintf("create volume %s success\n", newVolumeName2))
	}
}

func teardown() {
	fmt.Printf("delete source snapshot %s\n", srcSnapshotName)
	deleteSnapshotParameter := parameters.DeleteSnapshot{}
	deleteSnapshotParameter.SetVolumeName(srcVolumeName)
	deleteSnapshotParameter.SetSnapshotName(srcSnapshotName)

	err := fassSDK.DeleteSnapshot(&deleteSnapshotParameter, uuid.New().String())
	if err != nil {
		fmt.Printf("delete source snapshot %s failed due to %s\n", srcSnapshotName, err.Error())
	} else {
		fmt.Printf("delete source snapshot %s success\n", srcSnapshotName)
	}

	fmt.Printf("delete source subsys %s\n", srcSubsysName)
	deleteSubsysParameter := parameters.DeleteSubsys{}
	deleteSubsysParameter.SetSubsysName(srcSubsysName)
	deleteSubsysParameter.ForceDelete()
	deleteSubsysParameter.DeleteVolume()

	err = fassSDK.DeleteSubsys(&deleteSubsysParameter, uuid.New().String())
	if err != nil {
		fmt.Printf("delete source subsys %s failed due to %s\n", srcSubsysName, err.Error())
	} else {
		fmt.Printf("delete source subsys %s success\n", srcSubsysName)
	}

	err = deleteVolume(newVolumeName2, 3)
	if err != nil {
		fmt.Printf("delete volume %s failed due to %s\n", newVolumeName2, err.Error())
	} else {
		fmt.Printf("delete volume %s success\n", newVolumeName2)
	}
}

func createVolumeFromSnapshot(subsysName, volumeName string) error {
	createSubsysParameter := parameters.CreateSubsysFromSnapshot{}
	createSubsysParameter.SetName(subsysName)
	createSubsysParameter.SetPoolName(poolName)
	createSubsysParameter.SetVolumeName(volumeName)
	createSubsysParameter.SetSrcVolumeName(srcVolumeName)
	createSubsysParameter.SetSrcSnapshotName(srcSnapshotName)
	createSubsysParameter.EnableISCSI()

	_, err := fassSDK.CreateSubsysFromSnapshot(&createSubsysParameter, uuid.New().String())
	if err != nil {
		return err
	}

	deleteSubsysParameter := parameters.DeleteSubsys{}
	deleteSubsysParameter.SetSubsysName(subsysName)

	err = fassSDK.DeleteSubsys(&deleteSubsysParameter, uuid.New().String())
	if err != nil {
		return err
	}

	return nil
}

func TestCreateVolume(t *testing.T) {
	err := createVolumeFromSnapshot(newSubsysName1, newVolumeName1)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestListVolume(t *testing.T) {
	parameter := parameters.ListVolume{}

	_, err := fassSDK.ListVolume(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveVolume(t *testing.T) {
	parameter := parameters.RetrieveVolume{}
	parameter.SetVolumeName(newVolumeName1)

	_, err := fassSDK.RetrieveVolume(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveVolumeNotExists(t *testing.T) {
	parameter := parameters.RetrieveVolume{}
	parameter.SetVolumeName(invalidVolumeName)

	_, err := fassSDK.RetrieveVolume(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fe, ok := err.(*fassSDK.SDKError)
		if ok {
			if !fe.IsNotExists() {
				t.Fail()
				return
			}
			if !fe.IsExists() {
				t.Fail()
				return
			}
			return
		}
	}
	fmt.Printf("%s", err.Error())
	t.FailNow()
}

func TestExpandVolume(t *testing.T) {
	parameter := parameters.ExpandVolume{}
	parameter.SetVolumeName(newVolumeName1)
	parameter.SetNewCapacityGB(20)

	err := fassSDK.ExpandVolume(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestSetQosOfVolume(t *testing.T) {
	parameter := parameters.SetQosOfVolume{}
	parameter.SetVolumeName(newVolumeName1)
	parameter.SetBpsMB(100)
	parameter.SetIops(1000)

	err := fassSDK.SetQosOfVolume(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func flattenVolume(volumeName string) (*responses.FlattenVolume, error) {
	parameter := parameters.FlattenVolume{}
	parameter.SetVolumeName(volumeName)

	data, err := fassSDK.FlattenVolume(&parameter, uuid.New().String())
	return data, err
}

func TestFlattenVolume(t *testing.T) {
	data, err := flattenVolume(newVolumeName1)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}

	taskId = data.TaskId
	fmt.Printf("flatten volume task_id %s\n", taskId)
}

func TestFlattenVolumeProgress(t *testing.T) {
	parameter := parameters.GetFlattenVolumeProgress{}
	parameter.SetTaskId(taskId)

	for {
		data, err := fassSDK.FlattenVolumeProgress(&parameter, uuid.New().String())
		if !reflect.DeepEqual(err, nil) {
			fmt.Printf("%s", err.Error())
			t.FailNow()
		}
		if data.IsDone() {
			fmt.Printf("volume flatten complete\n")
			return
		}

		fmt.Printf("wait voluem faltten end\n")
		time.Sleep(2 * time.Second)
	}
}

func TestStopFlattenVolume(t *testing.T) {
	data, err := flattenVolume(newVolumeName2)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}

	parameter := parameters.StopFlattenVolume{}
	parameter.SetTaskId(data.TaskId)

	err = fassSDK.StopFlattenVolume(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func deleteVolume(volumeName string, retry int) error {
	parameter := parameters.DeleteVolume{}
	parameter.SetVolumeName(volumeName)
	parameter.ForceDelete()

	var err error
	err = fassSDK.DeleteVolume(&parameter, uuid.New().String())

	if retry == 0 {
		return err
	}

	for i := 1; i < retry; i++ {
		err = fassSDK.DeleteVolume(&parameter, uuid.New().String())
		if err == nil {
			return err
		}
		time.Sleep(3 * time.Second)
	}

	return err
}

func TestDeleteVolume(t *testing.T) {
	err := deleteVolume(newVolumeName1, 0)
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
