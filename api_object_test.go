package zabbix_test

import (
	zapi "github.com/claranet/go-zabbix-api"
	"reflect"
	"testing"
)

func testCRUDAPIObjectOperations(t *testing.T, object zapi.APIObject) {
	testCreateAPIObject(t, object)
	defer testDeleteAPIObject(t, object)
	reflectType := reflect.ValueOf(object).Elem().Type()
	emptyObject := reflect.New(reflectType).Interface().(zapi.APIObject)
	emptyObject.SetID(object.GetID())
	testReadAPIObject(t, emptyObject)
	testUpdateAPIObject(t, object)
}

func testCreateAPIObject(t *testing.T, object zapi.APIObject) {
	err := _api.CreateAPIObject(object)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Created API object: %+v", object)
}

func testReadAPIObject(t *testing.T, object zapi.APIObject) {
	err := _api.ReadAPIObject(object)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Read API object: %+v", object)
}

func testUpdateAPIObject(t *testing.T, object zapi.APIObject) {
	err := _api.UpdateAPIObject(object)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Updated API object: %+v", object)
}

func testDeleteAPIObject(t *testing.T, object zapi.APIObject) {
	err := _api.DeleteAPIObject(object)
	if err != nil {
		t.Fatal(err)
	}
	err = _api.ReadAPIObject(object)
	if err == nil {
		t.Fatal("Could not delete object")
	}
	t.Logf("Deleted API object: %+v", object)
}
