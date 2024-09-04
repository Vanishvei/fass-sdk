package account

// File       : account_test.go
// Path       : test/account
// Time       : CST 2023/5/5 14:55
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
	poolName        = "fast_pool"
	subsysName      = "s1000"
	volumeName      = "v1000"
	accountName     = "account_1"
	accountPassword = "admin@123456"
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

func TestCreateAccount(t *testing.T) {
	parameter := parameters.CreateAccount{}
	parameter.SetAccountName(accountName)
	parameter.SetPassword(accountPassword)
	resp, err := fassSDK.CreateAccount(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(resp.AccountName, accountName) {
		fmt.Printf("wrong account name")
		t.FailNow()
	}
}

func TestDuplicateCreateAccount(t *testing.T) {
	parameter := parameters.CreateAccount{}
	parameter.SetAccountName(accountName)
	parameter.SetPassword(accountPassword)
	_, err := fassSDK.CreateAccount(&parameter, uuid.New().String())
	if reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	var ex fassSDK.SDKError
	if errors.As(err, &ex) {
		if ex.IsExists() {
			return
		}
	}
	fmt.Printf("%s", err.Error())
	t.FailNow()

}

func TestListAccount(t *testing.T) {
	parameter := parameters.ListAccount{}
	resp, err := fassSDK.ListAccount(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(len(*resp), 1) {
		fmt.Printf("wrong account count")
		t.FailNow()
	}
}

func TestRetrieveAccount(t *testing.T) {
	parameter := parameters.RetrieveAccount{}
	parameter.SetAccountName(accountName)

	resp, err := fassSDK.RetrieveAccount(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(resp.AccountName, accountName) {
		fmt.Printf("account result %s is not equal to %s\n", resp.AccountName, accountName)
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

func TestDeleteAccountForbidden(t *testing.T) {
	parameter := parameters.DeleteAccount{}
	parameter.SetAccountName(accountName)
	err := fassSDK.DeleteAccount(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		var ex fassSDK.SDKError
		if errors.As(err, &ex) {
			if *ex.Code == 801001 {
				return
			}
		}
	}
	fmt.Printf("%s", err.Error())
	t.FailNow()
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

func TestDeleteAccount(t *testing.T) {
	parameter := parameters.DeleteAccount{}
	parameter.SetAccountName(accountName)
	err := fassSDK.DeleteAccount(&parameter, uuid.New().String())
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
