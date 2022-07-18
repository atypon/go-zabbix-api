package zabbix

// TemplateGroup represent Zabbix Template Group type returned from Zabbix API
type TemplateGroup struct {
	Name     string `json:"name"`
	GroupIDs string `json:"groupids,omitempty"`
	GroupID  string `json:"groupid,omitempty"`
}

// TemplateGroups  is an Array of Template structs.
type TemplateGroups []TemplateGroup

// TemplateGroupGet Wrapper for templategroup.get
func (api *API) TemplateGroupGet(params Params) (res TemplateGroups, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("templategroup.get", params, &res)
	return
}

// TemplateGroupCreate Wrapper for template.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/template/create
func (api *API) TemplateGroupCreate(TemplateGroups TemplateGroups) (err error) {
	response, err := api.CallWithError("templategroup.create", TemplateGroups)
	if err != nil {
		return
	}
	result := response.Result.(map[string]interface{})
	ids := result["groupids"].([]interface{})
	for i, id := range ids {
		TemplateGroups[i].GroupIDs = id.(string)
	}
	return
}

// TemplateGroupUpdate Wrapper for template.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/template/update
func (api *API) TemplateGroupUpdate(TemplateGroups TemplateGroups) (err error) {
	_, err = api.CallWithError("templategroup.update", TemplateGroups)
	return
}

// TemplateGroupDelete Wrapper for template.delete
// Cleans ApplicationID in all apps elements if call succeed.
// https://www.zabbix.com/documentation/3.2/manual/api/reference/template/delete
func (api *API) TemplateGroupDelete(TemplateGroups TemplateGroups) (err error) {
	TemplateGroupIds := make([]string, len(TemplateGroups))
	for i, template := range TemplateGroups {
		TemplateGroupIds[i] = template.GroupID
	}

	err = api.TemplateGroupDeleteByIds(TemplateGroupIds)
	if err == nil {
		for i := range TemplateGroups {
			TemplateGroups[i].GroupIDs = ""
		}
	}
	return
}

// TemplateGroupDeleteByIds Wrapper for template.delete
// Use template's id to delete the template
// https://www.zabbix.com/documentation/3.2/manual/api/reference/template/delete
func (api *API) TemplateGroupDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("templategroup.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	idss := result["groupids"].([]interface{})
	if len(ids) != len(idss) {
		err = &ExpectedMore{len(ids), len(idss)}
	}

	return
}
