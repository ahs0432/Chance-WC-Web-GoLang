package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var Config *ConfigData

func Init() {
	Config = configFileScan("config.yml")
	if Config == nil {
		os.Exit(3)
	}
}

func configFileCreate(confFile string) bool {
	databaseInfo := DatabaseInfo{
		Host:     "127.0.0.1",
		Port:     "3306",
		Protocol: "tcp",
		User:     "webchecker",
		Password: "P@ssW0rd",
		Name:     "webchecker"}

	configData := ConfigData{
		Port:         "80",
		DatabaseInfo: databaseInfo}

	data, err := yaml.Marshal(&configData)

	if err != nil {
		fmt.Println(confFile + " file create failed")
		fmt.Println(err)
		return false
	}

	err2 := ioutil.WriteFile(confFile, data, 0644)
	if err2 != nil {
		fmt.Println(confFile + " file create failed")
		fmt.Println(err)
		return false
	}

	fmt.Println(confFile + " file create")
	return true
}

func configFileScan(confFile string) *ConfigData {
	filename, _ := filepath.Abs(confFile)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), "The system cannot find the file specified.") || strings.Contains(err.Error(), "no such file or directory") {
			if !configFileCreate(confFile) {
				return nil
			}

			return configFileScan(confFile)
		} else {
			return nil
		}
	}

	configData := &ConfigData{}
	err = yaml.Unmarshal([]byte(yamlFile), &configData)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return configData
}

type ConfigData struct {
	Port         string       `yaml:"Port"`
	DatabaseInfo DatabaseInfo `yaml:"DatabaseInfo"`
}

type DatabaseInfo struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Protocol string `yaml:"Protocol"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Name     string `yaml:"Name"`
}
