package pool

// File       : pool_test.go
// Path       : requests
// Time       : CST 2023/4/25 17:28
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"

	fassSDK "github.com/Vanishvei/fass-sdk"
	parameters "github.com/Vanishvei/fass-sdk-parameters"
	_ "github.com/Vanishvei/fass-sdk-responses"
	_ "github.com/Vanishvei/fass-sdk/test"
)

var poolName = "fast_pool"
var invalidPoolName = "pool9999"

func TestRetrievePool(t *testing.T) {
	parameter := parameters.RetrievePool{}
	parameter.SetPoolName(poolName)
	_, err := fassSDK.RetrievePool(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrievePoolNotExists(t *testing.T) {
	parameter := parameters.RetrievePool{}
	parameter.SetPoolName(invalidPoolName)
	_, err := fassSDK.RetrievePool(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		var fe fassSDK.SDKError
		if errors.As(err, &fe) {
			if fe.IsNotExists() {
				return
			}
		}
	}
	fmt.Printf("%s", err.Error())
	t.FailNow()
}
