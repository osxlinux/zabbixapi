package utils

import (
	"fmt"
	"testing"
)

func TestNewConfigMgr(t *testing.T) {
	var err error
	if err = NewConfigMgr(`/Users/osx/Documents/go/src/zabbixapi/conf/zabbix.json`); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(GlobConfigMgr.API)
	fmt.Println(GlobConfigMgr.Password)
	fmt.Println(GlobConfigMgr.User)
}

func TestNewConfigMgr2(t *testing.T) {
	var err error
	if err = NewConfigMgr(`/Users/osx/Documents/go/src/zabbixapi/conf/zabbix.json`); err != nil {
		t.Error(err.Error())
	}

}