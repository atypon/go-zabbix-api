package zabbix_test

import (
	zapi "github.com/claranet/go-zabbix-api"
	"testing"
)

func testCRUDAPIObjectOperations(t *testing.T, object zapi.APIObject) {
	testCreateAPIObject(t, object)
	defer testDeleteAPIObject(t, object)
	testReadAPIObject(t, object)
	testUpdateAPIObject(t, object)
}

func testCreateAPIObject(t *testing.T, object zapi.APIObject) {
	err := _api.CreateAPIObject(object)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Created API object: %s", object)
}

func testReadAPIObject(t *testing.T, object zapi.APIObject) {
	err := _api.ReadAPIObject(object)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Read API object: %s", object)
}

func testUpdateAPIObject(t *testing.T, object zapi.APIObject) {
	err := _api.UpdateAPIObject(object)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Updated API object: %s", object)
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
	t.Logf("Deleted API object: %s", object)
}
