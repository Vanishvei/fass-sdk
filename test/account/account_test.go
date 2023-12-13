package account

// File       : account_test.go
// Path       : test/account
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
	accountName     = "account_1"
	accountPassword = "admin@123"
)

func TestCreateAccount(t *testing.T) {
	parameter := parameters.CreateAccount{}
	parameter.SetAccountName(accountName)
	parameter.SetPassword(accountPassword)
	_, err := fassSDK.CreateAccount(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestListAccount(t *testing.T) {
	parameter := parameters.ListAccount{}
	_, err := fassSDK.ListAccount(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
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

func TestDeleteAccount(t *testing.T) {
	parameter := parameters.DeleteAccount{}
	parameter.SetAccountName(accountName)
	err := fassSDK.DeleteAccount(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}
