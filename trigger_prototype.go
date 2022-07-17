package zabbix

// TriggerPrototype represent Zabbix trigger prototype object
// https://www.zabbix.com/documentation/3.2/manual/api/reference/triggerprototype/object
type TriggerPrototype struct {
	TriggerID          string            `json:"triggerid,omitempty"` // Readonly
	Description        string            `json:"description"`         // Reqired
	Expression         string            `json:"expression"`          // Required
	Commemts           string            `json:"comments,omitempty"`
	Priority           SeverityType      `json:"priority,omitempty,string"`
	Status             StatusType        `json:"status,omitempty,string"`
	TemplateID         string            `json:"templateid,omitempty"` // Readonly
	Type               int               `json:"type,omitempty,string"`
	URL                string            `json:"url,omitempty"`
	RecoveryMode       int               `json:"recovery_mode,omitempty,string"`
	RecoveryExpression string            `json:"recovery_expression,omitempty"`
	CorrelationMode    int               `json:"correlation_mode,omitempty,string"`
	CorrelationTag     string            `json:"correlation_tag,omitempty"`
	ManualClose        int               `json:"manual_close,omitempty,string"`
	Dependencies       TriggerPrototypes `json:"dependencies,omitempty"`

	Functions TriggerFunctions `json:"functions,omitempty"`
	// Return the hosts that the trigger prototype belongs to in the hosts property.
	ParentHosts Hosts `json:"hosts,omitempty"`
}

// TriggerPrototypes is an array of TriggerPrototype
type TriggerPrototypes []TriggerPrototype

// TriggerPrototypesGet Wrapper for trigger.get
// https://www.zabbix.com/documentation/3.2/manual/api/reference/triggerprototype/get
func (api *API) TriggerPrototypesGet(params Params) (res TriggerPrototypes, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("triggerprototype.get", params, &res)
	return
}

// TriggerPrototypeGetByID Gets trigger by Id only if there is exactly 1 matching trigger.
func (api *API) TriggerPrototypeGetByID(id string) (res *TriggerPrototype, err error) {
	triggers, err := api.TriggerPrototypesGet(Params{"triggerids": id})
	if err != nil {
		return
	}

	if len(triggers) != 1 {
		e := ExpectedOneResult(len(triggers))
		err = &e
		return
	}
	res = &triggers[0]
	return
}

// TriggerPrototypesCreate Wrapper for trigger.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/triggerprototype/create
func (api *API) TriggerPrototypesCreate(triggers TriggerPrototypes) (err error) {
	response, err := api.CallWithError("triggerprototype.create", triggers)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	triggerids := result["triggerids"].([]interface{})
	for i, id := range triggerids {
		triggers[i].TriggerID = id.(string)
	}
	return
}

// TriggerPrototypesUpdate Wrapper for trigger.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/triggerprototype/update
func (api *API) TriggerPrototypesUpdate(triggers TriggerPrototypes) (err error) {
	_, err = api.CallWithError("triggerprototype.update", triggers)
	return
}

// TriggerPrototypesDelete Wrapper for trigger.delete
// Cleans TriggerID in all triggers elements if call succeed.
// https://www.zabbix.com/documentation/3.2/manual/api/reference/triggerprototype/delete
func (api *API) TriggerPrototypesDelete(triggers TriggerPrototypes) (err error) {
	ids := make([]string, len(triggers))
	for i, trigger := range triggers {
		ids[i] = trigger.TriggerID
	}

	err = api.TriggerPrototypesDeleteByIds(ids)
	if err == nil {
		for i := range triggers {
			triggers[i].TriggerID = ""
		}
	}
	return
}

// TriggerPrototypesDeleteByIds Wrapper for trigger.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/triggerprototype/delete
func (api *API) TriggerPrototypesDeleteByIds(ids []string) (err error) {
	triggerids1, err := api.TriggerPrototypesDeleteIDs((ids))
	if err != nil {
		return
	}

	if len(triggerids1) != len(ids) {
		err = &ExpectedMore{len(ids), len(triggerids1)}
	}
	return
}

// TriggerPrototypesDeleteIDs Wrapper for trigger.delete
// return the id of the deleted trigger prototype
// https://www.zabbix.com/documentation/3.2/manual/api/reference/triggerprototype/delete
func (api *API) TriggerPrototypesDeleteIDs(ids []string) (triggerids []interface{}, err error) {
	response, err := api.CallWithError("triggerprototype.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	triggerids1, ok := result["triggerids"].([]interface{})
	if !ok {
		triggerids2 := result["triggerids"].(map[string]interface{})
		for _, id := range triggerids2 {
			triggerids = append(triggerids, id)
		}
	} else {
		triggerids = triggerids1
	}
	return
}
