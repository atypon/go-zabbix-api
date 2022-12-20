package zabbix

import (
	"encoding/json"
	"fmt"
	"strings"
)

type APIObject interface {
	GetID() string
	SetID(string)
	GetAPIModule() string
}

func (api *API) CreateAPIObject(object APIObject) (err error) {
	method := fmt.Sprintf("%s.create", object.GetAPIModule())
	response, err := api.CallWithError(method, object)
	if err != nil {
		return
	}
	result := response.Result.(map[string]any)
	id, err := getIDFromCreateResult(result)
	object.SetID(id)
	return
}

func getIDFromCreateResult(result map[string]any) (id string, err error) {
	for key, value := range result {
		if strings.HasSuffix(key, "ids") {
			ids, ok := value.([]any)
			if !ok || len(ids) == 0 {
				err = fmt.Errorf("couldn't find id: %s", ids)
				return
			}
			id = ids[0].(string)
			break
		}
	}
	return
}

func (api *API) ReadAPIObject(object APIObject) (err error) {
	var objects []json.RawMessage
	method := fmt.Sprintf("%s.get", object.GetAPIModule())
	err = api.CallWithErrorParse(method, Params{"filter": object}, &objects)
	if err != nil {
		return
	}
	if len(objects) == 0 {
		err = fmt.Errorf("%s with ID: %s not found", object.GetAPIModule(), object.GetID())
		return
	}
	err = json.Unmarshal(objects[0], &object)
	if err != nil {
		return
	}
	return
}

func (api *API) UpdateAPIObject(object APIObject) (err error) {
	method := fmt.Sprintf("%s.update", object.GetAPIModule())
	_, err = api.CallWithError(method, object)
	return
}

func (api *API) DeleteAPIObject(object APIObject) (err error) {
	method := fmt.Sprintf("%s.delete", object.GetAPIModule())
	_, err = api.CallWithError(method, []string{object.GetID()})
	return
}
