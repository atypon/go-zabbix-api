package zabbix

// ItemPrototype represent Zabbix item prototype object
// https://www.zabbix.com/documentation/3.2/manual/api/reference/itemprototype/object
type ItemPrototype struct {
	ItemID               string    `json:"itemid,omitempty"`  // Readonly
	Delay                string    `json:"delay"`             // Required
	HostID               string    `json:"hostid"`            // Required
	InterfaceID          string    `json:"interfaceid"`       // Required
	Key                  string    `json:"key_"`              // Required
	Name                 string    `json:"name"`              // Required
	Type                 ItemType  `json:"type,string"`       // Required
	ValueType            ValueType `json:"value_type,string"` // Required
	RuleID               string    `json:"ruleid"`            // Required for item prototype creation
	AuthType             int       `json:"authtype,omitempty,string"`
	DataType             DataType  `json:"data_type,omitempty,string"`
	DelayFlex            string    `json:"delay_flex,omitempty"`
	Delta                DeltaType `json:"delta,omitempty,string"`
	Description          string    `json:"description,omitempty"`
	History              string    `json:"history,omitempty"`
	IpmiSensor           string    `json:"ipmi_sensor,omitempty"`
	Logtimefmt           string    `json:"logtimefmt,omitempty"`
	Multiplier           string    `json:"multiplier,omitempty"`
	Params               string    `json:"params,omitempty"`
	Password             string    `json:"password,omitempty"`
	Port                 string    `json:"port,omitempty"`
	PrivateKey           string    `json:"privatekey,omitempty"`
	PublicKey            string    `json:"publickey,omitempty"`
	SnmpCommunity        string    `json:"snmp_community,omitempty"`
	SnmpOid              string    `json:"snmp_oid,omitempty"`
	Snmpv3Authpassphrase string    `json:"snmpv3_authpassphrase,omitempty"`
	Snmpv3Authprotocol   int       `json:"snmpv3_authprotocol,omitempty,string"`
	Snmpv3Contextname    string    `json:"snmpv3_contextname,omitempty"`
	Snmpv3Privpassphrase string    `json:"snmpv3_privpassphrase,omitempty"`
	Snmpv3Privprotocol   int       `json:"snmpv3_privprotocol,omitempty,string"`
	Snmpv3Securitylevel  int       `json:"snmpv3_securitylevel,omitempty,string"`
	Snmpv3Securityname   string    `json:"snmpv3_securityname,omitempty"`
	Status               int       `json:"status,string"`
	Templateid           string    `json:"templateid,omitempty"`
	TrapperHosts         string    `json:"trapper_hosts,omitempty"`
	Trends               string    `json:"trends,omitempty"`
	Units                string    `json:"units,omitempty"`
	Username             string    `json:"username,omitempty"`
	Valuemapid           string    `json:"valuemapid,omitempty"`

	DiscoveryRule LLDRule `json:"DiscoveryRule,omitempty"`
	Hosts         Hosts   `json:"hosts,omitempty"`
}

// ItemPrototypes is an array of ItemPrototype
type ItemPrototypes []ItemPrototype

// ItemPrototypesGet Wrapper for item.get
// https://www.zabbix.com/documentation/3.2/manual/api/reference/itemprototype/get
func (api *API) ItemPrototypesGet(params Params) (res ItemPrototypes, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("itemprototype.get", params, &res)
	return
}

// ItemPrototypeGetByID Gets item by Id only if there is exactly 1 matching item.
func (api *API) ItemPrototypeGetByID(id string) (res *ItemPrototype, err error) {
	items, err := api.ItemPrototypesGet(Params{"itemids": id})
	if err != nil {
		return
	}

	if len(items) != 1 {
		e := ExpectedOneResult(len(items))
		err = &e
		return
	}
	res = &items[0]
	return
}

// ItemPrototypesCreate Wrapper for item.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/itemprototype/create
func (api *API) ItemPrototypesCreate(items ItemPrototypes) (err error) {
	response, err := api.CallWithError("itemprototype.create", items)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	itemids := result["itemids"].([]interface{})
	for i, id := range itemids {
		items[i].ItemID = id.(string)
	}
	return
}

// ItemPrototypesUpdate Wrapper for item.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/itemprototype/update
func (api *API) ItemPrototypesUpdate(items ItemPrototypes) (err error) {
	_, err = api.CallWithError("itemprototype.update", items)
	return
}

// ItemPrototypesDelete Wrapper for item.delete
// Cleans ItemId in all items elements if call succeed.
// https://www.zabbix.com/documentation/3.2/manual/api/reference/itemprototype/delete
func (api *API) ItemPrototypesDelete(items ItemPrototypes) (err error) {
	ids := make([]string, len(items))
	for i, item := range items {
		ids[i] = item.ItemID
	}

	err = api.ItemPrototypesDeleteByIds(ids)
	if err == nil {
		for i := range items {
			items[i].ItemID = ""
		}
	}
	return
}

// ItemPrototypesDeleteByIds Wrapper for item.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/itemprototype/delete
func (api *API) ItemPrototypesDeleteByIds(ids []string) (err error) {
	itemids1, err := api.ItemPrototypesDeleteIDs(ids)
	if err != nil {
		return
	}

	if len(itemids1) != len(ids) {
		err = &ExpectedMore{len(ids), len(itemids1)}
	}
	return
}

// ItemPrototypesDeleteIDs Wrapper for item.delete
// Delete the item prototype and return the id of the deleted item prototype
// https://www.zabbix.com/documentation/3.2/manual/api/reference/itemprototype/delete
func (api *API) ItemPrototypesDeleteIDs(ids []string) (itemids1 []interface{}, err error) {
	response, err := api.CallWithError("itemprototype.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	itemids1 = result["prototypeids"].([]interface{})
	return
}
