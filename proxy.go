package zabbix

type Proxy struct {
	ProxyID        string               `json:"proxyid,omitempty" zabbix:"id"`
	Name           string               `json:"host"`
	MonitoredHosts []ProxyMonitoredHost `json:"monitored_hosts,omitempty"`
	Status         ProxyStatus          `json:"status,string"`
	Description    string               `json:"description,omitempty"`
	ProxyAddress   string               `json:"proxy_address,omitempty"`
	Interface      any                  `json:"interface,omitempty"`
}

type ProxyStatus int

type ProxyMonitoredHost struct {
	HostID string `json:"hostid"`
}

type ProxyInterface struct {
	//InterfaceID string `json:"interfaceid,omitempty"`
	DNS   string `json:"dns"`
	IP    string `json:"ip"`
	Port  string `json:"port"`
	UseIP int    `json:"useip,string"`
}

//func (p ProxyInterface) MarshalJSON() ([]byte, error) {
//	jsonMap := make(map[string]string)
//	if p.DNS != nil {
//		jsonMap["dns"] = *p.DNS
//	}
//	if p.IP != nil {
//		jsonMap["ip"] = *p.IP
//	}
//	jsonMap["port"] = p.Port
//	if p.UseIP == 0 {
//		jsonMap["useip"] = "0"
//	} else {
//		jsonMap["useip"] = "1"
//	}
//	return json.Marshal(jsonMap)
//}

const (
	ActiveProxy ProxyStatus = iota + 5
	PassiveProxy
)

func (p *Proxy) GetID() string {
	return p.ProxyID
}

func (p *Proxy) SetID(id string) {
	p.ProxyID = id
}

func (p *Proxy) GetAPIModule() string {
	return "proxy"
}

func (p *Proxy) GetExtraParams() Params {
	return Params{
		"selectHosts":     "extend",
		"selectInterface": "extend",
	}
}
