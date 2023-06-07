package test

// File       : init.go
// Path       : test
// Time       : CST 2023/4/28 15:18
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"

	fassSDK "github.com/Vanishvei/fass-sdk"
	"github.com/Vanishvei/fass-sdk/horizontal"
)

const suzakuApiConfigYaml = "test_config.yaml"

type suzakuApiConfig struct {
	EndpointList   []string `yaml:"fass_api_points"`
	Port           int      `yaml:"fass_api_port"`
	ConnectTimeout int      `yaml:"connect_timeout"`
	ReadTimeout    int      `yaml:"read_timeout"`
	Backoff        int      `yaml:"backoff"`
	RetryCount     int      `yaml:"retry_count"`
}

func init() {
	workDir, _ := os.Getwd()

	parentDir := horizontal.Substring(workDir, 0, strings.LastIndex(workDir, string(os.PathSeparator)))

	_configYaml := path.Join(parentDir, suzakuApiConfigYaml)
	_, err := os.Lstat(_configYaml)
	if os.IsNotExist(err) {
		panic(fmt.Sprintf("config yml file %s not exists", _configYaml))
	}

	data, err := ioutil.ReadFile(_configYaml)
	if err != nil {
		panic(fmt.Sprintf("Open config yaml failed. due to %s", err))
	}

	conf := suzakuApiConfig{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(fmt.Sprintf("Load config yaml failed. due to %s", err))
	}

	if len(conf.EndpointList) == 0 {
		panic("fass_api_points no config")
	}

	if conf.Port == 0 {
		conf.Port = 8000
	}

	if conf.ReadTimeout == 0 {
		conf.ReadTimeout = 300
	}

	if conf.Backoff == 0 {
		conf.Backoff = 3
	}

	if conf.ConnectTimeout == 0 {
		conf.ConnectTimeout = 1
	}

	fassSDK.InitConfig(
		&conf.EndpointList,
		&conf.Port,
		&conf.ReadTimeout,
		&conf.ConnectTimeout,
		&conf.Backoff,
		&conf.RetryCount)

}
