package cli

import (
	"fmt"
	"testing"
	"zabbixapi/src/utils"
)

func TestZabbixMgr_FetchSession(t *testing.T) {
	var err  error
	if err = InitHttpCliMgr(); err != nil {
		t.Error(err.Error())
	}
	if err = utils.NewConfigMgr(`/Users/osx/Documents/go/src/zabbixapi/conf/zabbix.json`); err != nil {
		t.Error(err.Error())
	}
	if err = InitZabbixMgr(); err != nil {
		t.Error(err.Error())
	}
	var s []ZabbixHosts //sessionid
	if s, err = GlobZabbixMgr.Hosts(); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(s)

}

func TestZabbixMgr_DefaultGroup(t *testing.T) {
	var err  error
	if err = InitHttpCliMgr(); err != nil {
		t.Error(err.Error())
	}
	if err = utils.NewConfigMgr(`/Users/osx/Documents/go/src/zabbixapi/conf/zabbix.json`); err != nil {
		t.Error(err.Error())
	}
	if err = InitZabbixMgr(); err != nil {
		t.Error(err.Error())
	}
	var s string //sessionid
	if s, err = GlobZabbixMgr.LinuxDefaultGroup(); err != nil {
		t.Error(err.Error())
	}

	fmt.Println(s)
}

func TestZabbixMgr_LinuxDefaultTemp(t *testing.T) {
	var err  error
	if err = InitHttpCliMgr(); err != nil {
		t.Error(err.Error())
	}
	if err = utils.NewConfigMgr(`/Users/osx/Documents/go/src/zabbixapi/conf/zabbix.json`); err != nil {
		t.Error(err.Error())
	}
	if err = InitZabbixMgr(); err != nil {
		t.Error(err.Error())
	}
	var s string //sessionid
	if s, err = GlobZabbixMgr.LinuxDefaultTemp(); err != nil {
		t.Error(err.Error())
	}

	fmt.Println(s)
}

func TestZabbixMgr_CreateHost(t *testing.T) {
	var err  error
	if err = InitHttpCliMgr(); err != nil {
		t.Error(err.Error())
	}
	if err = utils.NewConfigMgr(`/Users/osx/Documents/go/src/zabbixapi/conf/zabbix.json`); err != nil {
		t.Error(err.Error())
	}
	if err = InitZabbixMgr(); err != nil {
		t.Error(err.Error())
	}
	if err = GlobZabbixMgr.CreateHost(); err != nil {
		t.Error(err.Error())
	}
}