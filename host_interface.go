package zabbix

type (
	// InterfaceType different interface type
	InterfaceType int
)

const (
	// Differente type of zabbix interface
	// see "type" in https://www.zabbix.com/documentation/3.2/manual/api/reference/hostinterface/object

	// AgentInterface type
	AgentInterface InterfaceType = 1
	// SNMPInterface type
	SNMPInterface InterfaceType = 2
	// IPMIInterface type
	IPMIInterface InterfaceType = 3
	// JMXInterface type
	JMXInterface InterfaceType = 4
)

// HostInterface represents zabbix host interface type
type HostInterface struct {
	InterfaceID string        `json:"interfaceid,omitempty" zabbix:"id"`
	DNS         string        `json:"dns"`
	IP          string        `json:"ip"`
	Main        int           `json:"main,string"`
	Port        string        `json:"port"`
	Type        InterfaceType `json:"type,string"`
	UseIP       int           `json:"useip,string"`
	HostID      string        `json:"hostid,omitempty"`
	Details     any           `json:"details"`
}

type HostInterfaces []HostInterface

type SNMPDetails struct {
	Version        string `json:"version"`
	Bulk           string `json:"bulk,omitempty"`
	Community      string `json:"community,omitempty"`
	SecurityName   string `json:"securityname,omitempty"`
	SecurityLevel  string `json:"securitylevel,omitempty"`
	AuthPassphrase string `json:"authpassphrase,omitempty"`
	AuthProtocol   string `json:"authprotocol,omitempty"`
	PrivProtocol   string `json:"privprotocol,omitempty"`
	ContextName    string `json:"contextname,omitempty"`
}

func (hostInterface *HostInterface) GetID() string {
	return hostInterface.InterfaceID
}

func (hostInterface *HostInterface) SetID(id string) {
	hostInterface.InterfaceID = id
}

func (hostInterface *HostInterface) GetAPIModule() string {
	return "hostinterface"
}
