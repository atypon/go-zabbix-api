package zabbix

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type APIObject interface {
	GetID() string
	SetID(string)
	GetAPIModule() string
}

type ExtraParamsAPIObject interface {
	GetExtraParams() Params
}

func (api *API) CreateAPIObject(object APIObject) (err error) {
	method := fmt.Sprintf("%s.create", object.GetAPIModule())
	idsKey, err := getIDsKey(object)
	if err != nil {
		return
	}
	response, err := api.CallWithError(method, object)
	if err != nil {
		return
	}
	result := response.Result.(map[string]any)
	ids := result[idsKey].([]any)
	id := ids[0].(string)
	object.SetID(id)
	return
}

func (api *API) ReadAPIObject(object APIObject) (err error) {
	var objects []json.RawMessage
	method := fmt.Sprintf("%s.get", object.GetAPIModule())
	params, err := generateReadParams(object)
	if err != nil {
		return
	}
	err = api.CallWithErrorParse(method, params, &objects)
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

func generateReadParams(object APIObject) (params Params, err error) {
	_, ok := object.(ExtraParamsAPIObject)
	if ok {
		params = object.(ExtraParamsAPIObject).GetExtraParams()
	} else {
		params = Params{}
	}
	idsKey, err := getIDsKey(object)
	params[idsKey] = object.GetID()
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

func getIDsKey(object APIObject) (idsKey string, err error) {
	reflectType := reflect.ValueOf(object).Elem().Type()
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		if field.Tag.Get("zabbix") == "id" {
			jsonTag := field.Tag.Get("json")
			jsonKey := strings.Split(jsonTag, ",")[0]
			idsKey = fmt.Sprintf("%ss", jsonKey)
			return
		}
	}
	err = fmt.Errorf("%s type doesn't have a field with the tag zabbix:\"id\"", reflectType)
	return
}
