package zabbix

type MediaType struct {
	MediaID       string `json:"mediatypeid,omitempty"`
	MediaName     string `json:"name"`
	MediaType     string `json:"type"`
	EmailPassword string `json:"passwd"`
}

// Medias is an array of Media
type Medias []MediaType

func (api *API) MediaGet(params Params) (res Medias, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("mediatype.get", params, &res)
	return
}

// MacroGetByID Get macro by macro ID if there is exactly 1 matching macro
func (api *API) MediaGetByID(id string) (res *Macro, err error) {
	triggers, err := api.MacrosGet(Params{"hostmacroids": id})
	if err != nil {
		return
	}

	if len(triggers) == 1 {
		res = &triggers[0]
	} else {
		e := ExpectedOneResult(len(triggers))
		err = &e
	}
	return
}

// MacrosCreate Wrapper for usermacro.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/create
func (api *API) MediasCreate(macros Macros) error {
	response, err := api.CallWithError("mediatype.create", macros)
	if err != nil {
		return err
	}

	result := response.Result.(map[string]interface{})
	macroids := result["hostmacroids"].([]interface{})
	for i, id := range macroids {
		macros[i].HostID = id.(string)
	}
	return nil
}

// MacrosUpdate Wrapper for usermacro.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/update
func (api *API) MediasUpdate(macros Macros) (err error) {
	_, err = api.CallWithError("mediatype.create", macros)
	return
}

// MacrosDeleteByIDs Wrapper for usermacro.delete
// Cleans MacroId in all macro elements if call succeed.
//https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/delete
func (api *API) MediasDeleteByIDs(ids []string) (err error) {
	response, err := api.CallWithError("mediatype.delete", ids)

	result := response.Result.(map[string]interface{})
	hostmacroids := result["hostmacroids"].([]interface{})
	if len(ids) != len(hostmacroids) {
		err = &ExpectedMore{len(ids), len(hostmacroids)}
	}
	return
}

// MacrosDelete Wrapper for usermacro.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/delete
func (api *API) MediasDelete(macros Macros) (err error) {
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
