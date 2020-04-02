package cli

import (
	"encoding/json"
	"fmt"
	"strings"
	"zabbixapi/src/utils"
)

type ZabbixMgr struct {
	//session string `json:"session"`
}

var GlobZabbixMgr *ZabbixMgr

func InitZabbixMgr() (err error) {

	GlobZabbixMgr = &ZabbixMgr{}
	return
}

//获取Session TODO: 增加缓存
func (za *ZabbixMgr) session() (sessionid string, err error) {

	var requestBody ZabbixUserRequest
	var responseBody ZabbixUserResponse
	var pdata []byte
	requestBody.Jsonrpc = `2.0`
	requestBody.Method = `user.login`
	requestBody.Params.User = utils.GlobConfigMgr.User
	requestBody.Params.Password = utils.GlobConfigMgr.Password
	requestBody.Params.UserData = true
	requestBody.Id = 1

	if pdata, err = json.Marshal(requestBody); err != nil {
		return
	}

	var body []byte
	if body, err = GlobHttpCliMgr.RequestWithUrl("POST", utils.GlobConfigMgr.API, pdata); err != nil {
		return
	}

	if err = json.Unmarshal(body, &responseBody); err != nil {
		return
	}

	sessionid = responseBody.Result.Sessionid
	return
}

//获取主机列表
func (za *ZabbixMgr) Hosts() (hosts []ZabbixHosts, err error) {
	var zabbixHostsRequest ZabbixHostsRequest
	var zabbixHostsResponse ZabbixHostsResponse
	var zabbixHosts ZabbixHosts

	var pdata []byte
	var sessionid string
	if sessionid, err = za.session(); err != nil {
		return
	}

	zabbixHostsRequest.Id = 1
	zabbixHostsRequest.Jsonrpc = `2.0`
	zabbixHostsRequest.Method = `host.get`
	zabbixHostsRequest.Auth = sessionid
	zabbixHostsRequest.Params.Filter.Host = []string{`Zabbix Server`}

	if pdata, err = json.Marshal(zabbixHostsRequest); err != nil {
		return
	}
	var body []byte
	if body, err = GlobHttpCliMgr.RequestWithUrl(`POST`, utils.GlobConfigMgr.API, pdata); err != nil {
		return
	}
	if err = json.Unmarshal(body, &zabbixHostsResponse); err != nil {
		return
	}
	for _, v := range zabbixHostsResponse.Result {
		fmt.Println(v.Host, v.Hostid, v.Name)
		zabbixHosts.Name = v.Name
		zabbixHosts.Hostid = v.Hostid
		zabbixHosts.Host = v.Host
		hosts = append(hosts, zabbixHosts)
	}

	return
}

//创建一个主机
func (za *ZabbixMgr) CreateHost(hostname, tid , gid , ip string) (err error) {
	var (
		hcr      ZabbixCreateHostRequest
		session  string
		groupid  GroupId
		groupids []GroupId

		template  Template
		templates []Template

		interfaces  Interfaces
		interfacess []Interfaces

		pdata []byte
	)
	if session, err = za.session(); err != nil {
		return
	}

	hcr.Id = 1
	hcr.Auth = session
	hcr.Method = `host.create`
	hcr.Jsonrpc = `2.0`
	hcr.Params.InventoryMode = -1
	hcr.Params.Host = hostname

	//add group id
	//if groupid.Groupid, err = za.LinuxDefaultGroup(); err != nil {
	//	return
	//}
	groupid.Groupid = gid
	groupids = append(groupids, groupid)
	hcr.Params.Groups = groupids
	//add template id
	//if template.Templateid, err = za.LinuxDefaultTemp(); err != nil {
	//	return
	//}
	template.Templateid = tid
	templates = append(templates, template)
	hcr.Params.Templates = templates

	interfaces.Main = 1
	interfaces.Port = "10050"
	interfaces.Type = 1
	interfaces.UseIp = 1
	interfaces.Dns = ""
	interfaces.Ip = ip
	interfacess = append(interfacess, interfaces)
	hcr.Params.Interfaces = interfacess

	if pdata, err = json.Marshal(hcr); err != nil {
		return
	}
	fmt.Println(string(pdata))
	var body []byte
	if body, err = GlobHttpCliMgr.RequestWithUrl(`POST`, utils.GlobConfigMgr.API, pdata); err != nil {
		return
	}

	fmt.Println(string(body))

	return
}

//获取一个默认的主机组 Linux
func (za *ZabbixMgr) LinuxDefaultGroup() (groupid string, err error) {
	var (
		body              []byte
		hostGroupResponse ZabbixHostGroupResponse
	)
	if body, err = za.tgmethod(`hostgroup.get`); err != nil {
		return
	}
	if err = json.Unmarshal(body, &hostGroupResponse); err != nil {
		return
	}

	for _, result := range hostGroupResponse.Result {
		if strings.Contains(result.Name, `Linux servers`) {
			groupid = result.GroupId
			return
		}
	}
	return
}

//获取一个默认的模板
func (za *ZabbixMgr) LinuxDefaultTemp() (tid string, err error) {

	var (
		body     []byte
		response ZabbixTemplateResponse
	)
	if body, err = za.tgmethod(`template.get`); err != nil {
		return
	}
	if err = json.Unmarshal(body, &response); err != nil {
		return
	}
	for _, tmpr := range response.Result {
		if strings.Contains(tmpr.Name, `Template OS Linux`) {
			tid = tmpr.TemplateId
			return
		}
	}

	return
}

//command functiom
func (za *ZabbixMgr) tgmethod(m string) (body []byte, err error) {

	var (
		pdata     []byte
		sessionid string
	)

	if sessionid, err = za.session(); err != nil {
		return
	}
	zpub := zarepub{`2.0`, m, 1}
	var hostgroup = ZabbixHostGroupRequest{
		Params: HostGroupParams{
			Output: `extend`,
			Filter: HostGroupFilter{Name: []string{
				`Zabbix servers`,
				`Linux servers`,
				`Template OS Linux`,
				`Template OS Windows`,
			}},
		},
		Auth: sessionid,
	}
	hostgroup.zarepub = zpub
	fmt.Println(hostgroup)
	if pdata, err = json.Marshal(hostgroup); err != nil {
		return
	}

	if body, err = GlobHttpCliMgr.RequestWithUrl(`POST`, utils.GlobConfigMgr.API, pdata); err != nil {
		return
	}

	return
}
