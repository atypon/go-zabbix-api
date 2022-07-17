package zabbix_test

import (
	zapi "github.com/claranet/go-zabbix-api"
	"testing"
)

func TestAuth(t *testing.T) {
	api, err := testGetAPI(t).AuthGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Auth Result %s", api.Result)

}

func TestSetAuth(t *testing.T) {
	api, err := testGetAPI(t).AuthSet(zapi.Params{
		"authentication_type": 0,
		"http_auth_enabled":   0,
		"http_case_sensitive": 1})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Auth Setting Result %s", api.Result)

}
