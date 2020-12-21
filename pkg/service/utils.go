package service

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	PGUser       string  `json:"pguser"`
	PGPass       string  `json:"pgpass"`
	PGName       string  `json:"pgname"`
	PGHost       string  `json:"pghost"`
	PGPort       string  `json:"pgport"`
	VKToken      string  `json:"vktoken"`
	Timeout      float32 `json:"timeout"`
	VKApiVersion string  `json:"vkapi_version"`
}

func GetConfig(configPath string) (Config, error) {
	var config Config
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(bytes, &config)
	return config, err
}

func GetGroupsScreenNames(groupsPath string) ([]string, error) {
	bytes, err := ioutil.ReadFile(groupsPath)
	if err != nil {
		return nil, err
	}
	var groupsScreenNames []string
	err = json.Unmarshal(bytes, &groupsScreenNames)
	return groupsScreenNames, err
}

func GetUsersIDs(usersIDsPath string) ([]int32, error) {
	bytes, err := ioutil.ReadFile(usersIDsPath)
	if err != nil {
		return nil, err
	}
	var usersIDs []int32
	err = json.Unmarshal(bytes, &usersIDs)
	return usersIDs, err
}
