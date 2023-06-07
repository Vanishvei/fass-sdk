package acl

// File       : acl_test.go
// Path       : test/acl
// Time       : CST 2023/5/5 14:55
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
	_ "github.com/Vanishvei/fass-sdk-responses"
	_ "github.com/Vanishvei/fass-sdk/test"
)

var (
	hostName        = "client_1"
	groupName       = "group_1"
	accountName     = "account_1"
	accountPassword = "admin@123"

	addQualifierList = []string{
		"nqn.2019-06.suzaku:subsys001",
		"iqn.2019-03.cn.suzaku:subsys001",
		"iqn.2019-03.cn.suzaku:subsys002",
	}
	removeQualifierList = []string{
		"nqn.2019-06.suzaku:subsys001",
		"iqn.2019-03.cn.suzaku:subsys001",
	}
)

func TestCreateAccount(t *testing.T) {
	parameter := parameters.CreateAccountParameter{}
	parameter.SetAccountName(accountName)
	parameter.SetPassword(accountPassword)
	_, err := fassSDK.CreateAccount(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestListAccount(t *testing.T) {
	parameter := parameters.ListAccountParameter{}
	_, err := fassSDK.ListAccount(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveAccount(t *testing.T) {
	parameter := parameters.RetrieveAccountParameter{}
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

func TestDeleteAccount(t *testing.T) {
	parameter := parameters.DeleteAccountParameter{}
	parameter.SetAccountName(accountName)
	err := fassSDK.DeleteAccount(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestAddQualifierToGroup(t *testing.T) {
	parameter := parameters.AddQualifierToGroupParameter{}
	parameter.SetGroupName(groupName)
	parameter.SetHostName(hostName)
	parameter.SetQualifierList(addQualifierList)
	_, err := fassSDK.AddQualifierToGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRemoveQualifierFromGroup(t *testing.T) {
	parameter := parameters.RemoveQualifierFromGroupParameter{}
	parameter.SetGroupName(groupName)
	parameter.SetHostName(hostName)
	parameter.SetQualifierList(removeQualifierList)
	_, err := fassSDK.RemoveQualifierFromGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveGroup(t *testing.T) {
	parameter := parameters.RetrieveGroupParameter{}
	parameter.SetGroupName(groupName)

	resp, err := fassSDK.RetrieveGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(resp.GroupName, groupName) {
		fmt.Printf("group result %s is not equal to %s\n", resp.GroupName, groupName)
		t.FailNow()
	}
}

func TestListGroup(t *testing.T) {
	parameter := parameters.ListGroupParameter{}
	_, err := fassSDK.ListGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestDeleteGroup(t *testing.T) {
	parameter := parameters.DeleteGroupParameter{}
	parameter.SetGroupName(groupName)
	err := fassSDK.DeleteGroup(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}
