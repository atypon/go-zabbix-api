package zabbix

type (
	// Type of the event.
	// "source" in https://www.zabbix.com/documentation/3.2/manual/api/reference/event/object#event
	EventType string
)

const (
	// event created by a trigger
	TriggerEvent EventType = "0"
	// event created by a discovery rule
	DiscoveryRuleEvent EventType = "1"
	// event created by active agent auto-registration
	AutoRegistrationEvent EventType = "2"
	// internal event
	InternalEvent EventType = "3"
)
