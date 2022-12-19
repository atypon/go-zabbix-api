package zabbix

import (
	"encoding/json"
	"fmt"
)

type APIObject interface {
	GetID() string
	SetID(string)
	GetAPIModule() string
}

func (api *API) CreateAPIObject(object APIObject) (err error) {
	method := fmt.Sprintf("%s.create", object.GetAPIModule())
	idsKey := fmt.Sprintf("%sids", object.GetAPIModule())
	response, err := api.CallWithError(method, object)
	if err != nil {
		return
	}
	result := response.Result.(map[string]interface{})
	ids := result[idsKey].([]interface{})
	if len(ids) == 0 {
		return fmt.Errorf("could not create object")
	}
	object.SetID(ids[0].(string))
	return
}

func (api *API) ReadAPIObject(object APIObject) (err error) {
	var objects []json.RawMessage
	method := fmt.Sprintf("%s.get", object.GetAPIModule())
	idsKey := fmt.Sprintf("%sids", object.GetAPIModule())
	err = api.CallWithErrorParse(method, Params{idsKey: object.GetID()}, &objects)
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
