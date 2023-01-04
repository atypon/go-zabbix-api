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
