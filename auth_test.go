package zabbix_test

import (
	"encoding/json"
	"testing"
)

func TestAuth(t *testing.T) {
	authenticationSettings, err := _api.AuthGet()
	if err != nil {
		t.Fatal(err)
	}
	jsonObject, err := json.Marshal(authenticationSettings)
	t.Logf("auth get result %s", jsonObject)
	err = _api.AuthSet(authenticationSettings)
	if err != nil {
		t.Fatal(err)
	}
}
