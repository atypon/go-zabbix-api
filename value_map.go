package zabbix

// HostID in value map is the template ID to link to

type ValueMapType struct {
	ValueMapID       string       `json:"valuemapid,omitempty"`
	ValueMapName     string       `json:"name"`
	ValueMapMappings MappingsList `json:"mappings"`
	ValueMapHostID   string       `json:"hostid"`
	ValueMapUUID     string       `json:"uuid,omitempty"`
}
type Mapping struct {
	Type     int    `json:"type,omitempty"`
	Value    string `json:"value"`
	Newvalue string `json:"newvalue"`
}

type MappingsList []Mapping

// ValueMaps is an array of ValueMapType
type ValueMaps []ValueMapType

func (api *API) ValueMapGet(params Params) (res Medias, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("valuemap.get", params, &res)
	return
}

//// MacroGetByID Get macro by macro ID if there is exactly 1 matching macro
//func (api *API) ValueMapByID(id string) (res *Macro, err error) {
//	triggers, err := api.MacrosGet(Params{"hostmacroids": id})
//	if err != nil {
//		return
//	}
//
//	if len(triggers) == 1 {
//		res = &triggers[0]
//	} else {
//		e := ExpectedOneResult(len(triggers))
//		err = &e
//	}
//	return
//}

// MacrosCreate Wrapper for usermacro.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/create
func (api *API) ValueMapCreate(valueMaps ValueMaps) error {
	response, err := api.CallWithError("valuemap.create", valueMaps)
	if err != nil {
		return err
	}

	result := response.Result.(map[string]interface{})
	valuemapids := result["valuemapids"].([]interface{})

	for i, id := range valuemapids {
		valueMaps[i].ValueMapID = id.(string)
	}
	return nil
}

// MacrosUpdate Wrapper for usermacro.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/update
func (api *API) ValueMapUpdate(macros Macros) (err error) {
	_, err = api.CallWithError("valuemap.update", macros)
	return
}

//
//// MacrosDeleteByIDs Wrapper for usermacro.delete
//// Cleans MacroId in all macro elements if call succeed.
////https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/delete
//func (api *API) ValueMapDeleteByIDs(ids []string) (err error) {
//	response, err := api.CallWithError("valuemap.delete", ids)
//
//	result := response.Result.(map[string]interface{})
//	hostmacroids := result["hostmacroids"].([]interface{})
//	if len(ids) != len(hostmacroids) {
//		err = &ExpectedMore{len(ids), len(hostmacroids)}
//	}
//	return
//}

// MacrosDelete Wrapper for usermacro.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/delete
func (api *API) ValueMapDelete(macros Macros) (err error) {
	ids := make([]string, len(macros))
	for i, macro := range macros {
		ids[i] = macro.MacroID
	}

	err = api.MacrosDeleteByIDs(ids)
	if err == nil {
		for i := range macros {
			macros[i].MacroID = ""
		}
	}
	return
}
