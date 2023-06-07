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
	"strings"
)

type Config struct {
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

func (c *Config) SwitchEndpoint() {
	for _, endpoint := range *GlobalConfig.EndpointList {
		version, qps, err := getServerInfo(endpoint)
		if err != nil {
			continue
		}

		GlobalConfig.ApiQPS = &qps
		GlobalConfig.ApiVersion = &version
		GlobalConfig.CurrentEndpoint = &endpoint
		return
	}

	panic("Switch endpoint failed due to no normal nodes are available")
}

var GlobalConfig Config

func InitConfig(endpointList *[]string, port, readTimeout, connectTimeout, backoff, retryCount *int) {
	GlobalConfig.EndpointList = endpointList
	GlobalConfig.Port = port
	GlobalConfig.Backoff = backoff
	GlobalConfig.RetryCount = retryCount
	GlobalConfig.ReadTimeout = readTimeout
	GlobalConfig.ConnectTimeout = connectTimeout

	initServerInfo()
}

func initServerInfo() {
	for _, endpoint := range *GlobalConfig.EndpointList {
		version, qps, err := getServerInfo(endpoint)
		if err != nil {
			continue
		}

		if !strings.HasPrefix(version, "3.1") {
			panic(fmt.Sprintf("unsupported endpoint %s api version %s", endpoint, version))
		}

		GlobalConfig.ApiQPS = &qps
		GlobalConfig.ApiVersion = &version
		GlobalConfig.CurrentEndpoint = &endpoint
		return
	}

	if GlobalConfig.CurrentEndpoint == nil {
		panic("Init server info failed")
	}
}

func getServerInfo(endpoint string) (version string, qps int, err error) {
	response, err := http.Get(fmt.Sprintf("http://%s:%d/api/info", endpoint, *GlobalConfig.Port))
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
	return _serverInfo.Version, _serverInfo.ApiQps, nil
}

type serverInfo struct {
	Version     string `json:"version"`
	BuildDate   string `json:"build_date"`
	DeployModel string `json:"deploy_model"`
	WorkModel   string `json:"work_model"`
	Time        string `json:"time"`
	ApiQps      int    `json:"api_qps"`
}
