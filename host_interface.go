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
	Bulk           int    `json:"bulk,string,omitempty"`
	Community      string `json:"community"`
	SecurityName   string `json:"securityname"`
	SecurityLevel  int    `json:"securitylevel,string,omitempty"`
	AuthPassphrase string `json:"authpassphrase"`
	AuthProtocol   int    `json:"authprotocol,string,omitempty"`
	PrivProtocol   int    `json:"privprotocol,string,omitempty"`
	ContextName    string `json:"contextname"`
}

type SNMP3SecurityLevel int

type SNMP3AuthProtocol int

type SNMP3PrivProtocol int

const (
	SNMP3NoAuthNoPriv SNMP3SecurityLevel = iota
	SNMP3AuthNoPriv
	SNMP3AuthPriv
)

const (
	SNMP3MD5Auth SNMP3AuthProtocol = iota
	SNMP3SHA1Auth
	SNMP3SHA224Auth
	SNMP3SHA256Auth
	SNMP3SHA384Auth
	SNMP3SHA512Auth
)

const (
	SNMP3DESPriv SNMP3PrivProtocol = iota
	SNMP3AES128Priv
	SNMP3AES192Priv
	SNMP3AES256Priv
	SNMP3AES192CPriv
	SNMP3AES256CPriv
)

func (hostInterface *HostInterface) GetID() string {
	return hostInterface.InterfaceID
}

func (hostInterface *HostInterface) SetID(id string) {
	hostInterface.InterfaceID = id
}

func (hostInterface *HostInterface) GetAPIModule() string {
	return "hostinterface"
}
