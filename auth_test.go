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
	api, err := testGetAPI(t).AuthSet(zapi.AuthPrototype{
		HttpAuthEnabled:   0,
		HttpCaseSensitive: "1",
		LdapConfigured:    "0"})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Auth Setting Result %s", api.Result)

}
