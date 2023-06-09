package requests

// File       : config.go
// Path       : client
// Time       : CST 2023/4/24 16:00
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var allowVersionSet = map[string]string{"3.0": "v3"}

type config struct {
	Port            *int      `json:"port"`
	ConnectTimeout  *int      `json:"connectTimeout"`
	ReadTimeout     *int      `json:"readTimeout"`
	Backoff         *int      `json:"backoff"`
	RetryCount      *int      `json:"retryCount"`
	EndpointList    *[]string `json:"endpointList"`
	CurrentEndpoint *string   `json:"currentEndpoint"`
	ApiVersion      *string   `json:"api_version"`
	ApiQPS          *int      `json:"api_qps"`
}

func (c *config) SwitchEndpoint() {
	for _, endpoint := range *globalConfig.EndpointList {
		version, qps, err := getServerInfo(endpoint)
		if err != nil {
			continue
		}

		globalConfig.ApiQPS = &qps
		globalConfig.ApiVersion = &version
		globalConfig.CurrentEndpoint = &endpoint
		return
	}

	panic("Switch endpoint failed due to no normal nodes are available")
}

var globalConfig config

func InitConfig(endpointList *[]string, port, readTimeout, connectTimeout, backoff, retryCount *int) {
	globalConfig.EndpointList = endpointList
	globalConfig.Port = port
	globalConfig.Backoff = backoff
	globalConfig.RetryCount = retryCount
	globalConfig.ReadTimeout = readTimeout
	globalConfig.ConnectTimeout = connectTimeout

	initServerInfo()
}

func initServerInfo() {
	for _, endpoint := range *globalConfig.EndpointList {
		version, qps, err := getServerInfo(endpoint)
		if err != nil {
			continue
		}

		globalConfig.ApiQPS = &qps
		globalConfig.ApiVersion = &version
		globalConfig.CurrentEndpoint = &endpoint
		return
	}

	if globalConfig.CurrentEndpoint == nil {
		panic("Init server info failed")
	}
}

func getServerInfo(endpoint string) (version string, qps int, err error) {
	response, err := http.Get(fmt.Sprintf("http://%s:%d/api/info", endpoint, *globalConfig.Port))
	if err != nil {
		return version, qps, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return version, qps, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	var _serverInfo serverInfo
	err = json.Unmarshal(body, &_serverInfo)
	if err != nil {
		return version, qps, err
	}

	_ = response.Body.Close()
	return _serverInfo.APIVersion(endpoint), _serverInfo.ApiQps, nil
}

type serverInfo struct {
	Version     string `json:"version"`
	BuildDate   string `json:"build_date"`
	DeployModel string `json:"deploy_model"`
	WorkModel   string `json:"work_model"`
	Time        string `json:"time"`
	ApiQps      int    `json:"api_qps"`
}

func (s serverInfo) APIVersion(endpoint string) string {
	if len(s.Version) == 0 {
		panic("Get API version failed")
	}

	versionSectorCount := 3
	_version := strings.Split(s.Version, ".")
	if len(_version) != versionSectorCount {
		panic(fmt.Sprintf("Endpoint %s invalid API version %s", endpoint, s.Version))
	}

	for _, index := range []int{0, 1, 2} {
		_, err := strconv.Atoi(_version[index])
		if err != nil {
			panic(fmt.Sprintf("Endpoint %s invalid API version %s", endpoint, s.Version))
		}
	}

	apiVersion, ok := allowVersionSet[s.Version[0:3]]
	if !ok {
		panic(fmt.Sprintf("Unsupported endpoint %s api version %s", endpoint, s.Version))
	}

	return apiVersion
}
