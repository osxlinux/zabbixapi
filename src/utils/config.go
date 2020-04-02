package utils

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigMgr struct {
	API      string `json:"api" name:"api"`
	User     string `json:"user" name:"user"`
	Password string `json:"password" name:"password"`
}

var GlobConfigMgr *ConfigMgr

func NewConfigMgr(c string) (err error) {
	var cbyte []byte
	var config ConfigMgr
	if cbyte, err = ioutil.ReadFile(c); err != nil {
		return
	}

	if err = json.Unmarshal(cbyte, &config); err != nil {
		return
	}
	GlobConfigMgr = &config

	return
}
