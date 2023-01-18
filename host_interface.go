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
// https://www.zabbix.com/documentation/3.2/manual/api/reference/hostinterface/object
type HostInterface struct {
	InterfaceID string        `json:"interfaceid,omitempty"`
	DNS         string        `json:"dns"`
	IP          string        `json:"ip"`
	Main        int           `json:"main,string"`
	Port        string        `json:"port"`
	Type        InterfaceType `json:"type,string"`
	UseIP       int           `json:"useip,string"`
}

// HostInterfaces is an array of HostInterface
type HostInterfaces []HostInterface
