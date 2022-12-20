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
	filter, err := getIDReadFilter(object)
	if err != nil {
		return
	}
	err = api.CallWithErrorParse(method, Params{"filter": filter}, &objects)
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

// Returns a filter as a map of the form:
//
//	{
//			"objectids": object.GetID()
//	}
//
// Used because every zabbix api object has different
// key for the id
func getIDReadFilter(object APIObject) (filter map[string]any, err error) {
	filter = make(map[string]any)
	var jsonObject map[string]any
	data, err := json.Marshal(object)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &jsonObject)
	if err != nil {
		return
	}
	for key, value := range jsonObject {
		if value == object.GetID() {
			filter[key] = value
		}
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
