package pool

// File       : pool_test.go
// Path       : requests
// Time       : CST 2023/4/25 17:28
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

var poolName = "pool3"

func TestRetrievePool(t *testing.T) {
	parameter := parameters.RetrievePoolParameter{}
	parameter.SetPoolName(poolName)
	_, err := fassSDK.RetrievePool(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}
