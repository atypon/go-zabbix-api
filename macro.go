package zabbix

// Macro represent Zabbix User MAcro object
// https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/object
type Macro struct {
	MacroID   string `json:"hostmacroid,omitempty" zabbix:"id"`
	HostID    string `json:"hostid,omitempty"`
	MacroName string `json:"macro"`
	Value     string `json:"value"`
}

// Macros is an array of Macro
type Macros []Macro

func (m *Macro) GetID() string {
	return m.MacroID
}

func (m *Macro) SetID(id string) {
	m.MacroID = id
}

func (m *Macro) GetAPIModule() string {
	return "usermacro"
}
