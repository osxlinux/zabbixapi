package cli

const (
	LINUX_DEFAULT_TEMPLATE_ID   = `` //
	WINDOWS_DEFAULT_TEMPLATE_ID = ``
)

// Zabbix public request information
type zarepub struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Id      int    `json:"id"`
}

// zabbix public response information
type zareqpub struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
}

type ZabbixUserRequest struct {
	zarepub
	Params struct {
		User     string `json:"user"`
		Password string `json:"password"`
		UserData bool   `json:"userData"`
	} `json:"params"`
}

type ZabbixUserResponse struct {
	zareqpub
	Result struct {
		Sessionid string `json:"sessionid"`
	} `json:"result"`
}

type ZabbixHostsRequest struct {
	zarepub
	Auth   string `json:"auth"`
	Params struct {
		Filter struct {
			Host []string `json:"host"`
		}
	} `json:"params"`
}

type ZabbixHostsResponse struct {
	zareqpub
	Result []struct {
		Hostid string `json:"hostid"`
		Host   string `json:"host"`
		Name   string `json:"name"`
	} `json:"result" name:"result"`
}

type ZabbixHosts struct {
	Hostid string `json:"hostid"`
	Name   string `json:"name"`
	Host   string `json:"host"`
}

//Host Group
type ZabbixHostGroupRequest struct {
	zarepub
	Auth string `json:"auth"`

	Params HostGroupParams `json:"params"`
}

type HostGroupParams struct {
	Output string          `json:"output"`
	Filter HostGroupFilter `json:"filter"`
}
type HostGroupFilter struct {
	Name []string `json:"name"`
}

type ZabbixHostGroupResponse struct {
	zareqpub
	Result []struct {
		GroupId string `json:"groupid"`
		Name    string `json:"name"`
	}
}

type ZabbixHostCreateRequest struct {
	zarepub
	Auth   string `json:"auth"`
	Params struct {
	} `json:"params"`
}

type ZabbixTemplateResponse struct {
	zarepub
	Result []struct {
		TemplateId string `json:"templateid" name:"templateid"`
		Name       string `json:"name"`
	}
}

type ZabbixCreateHostRequest struct {
	zarepub
	Auth   string `json:"auth"`
	Params struct {
		Host          string       `json:"host"`
		Groups        []GroupId    `json:"groups"`
		Templates     []Template   `json:"templates"`
		Interfaces    []Interfaces `json:"interfaces"`
		InventoryMode int          `json:"inventory_mode"` //0 manual -1 disable 1 auto

	} `json:"params"`
}

type GroupId struct {
	Groupid string `json:"groupid"`
}

type Template struct {
	Templateid string `json:"templateid"`
}

type Interfaces struct {
	Ip    string `json:"ip"`
	Dns   string `json:"dns"`
	Port  string `json:"port"`
	Main  int    `json:"main"`
	Type  int    `json:"type"`
	UseIp int    `json:"useip"`
}
