package zabbix

func (api *API) AuthGet(params Params) (res Response, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	res, err = api.CallWithError("authentication.get", params)
	return
}

func (api *API) AuthSet(params Params) (res Response, err error) {
	res, err = api.CallWithError("authentication.update", params)
	return
}
