package zabbix_test

import (
	"testing"

	zapi "github.com/claranet/go-zabbix-api"
)

func TestUserGroupsGet(t *testing.T) {
	api := testGetAPI(t)

	params := zapi.Params{
		"filter": map[string]interface{}{
			"name": "Zabbix administrators",
		},
	}
	userGroups, err := api.UserGroupsGet(params)
	if err != nil {
		t.Fatal(err)
	}
	if len(userGroups) != 1 {
		t.Errorf("Bad user groups: %#v", userGroups)
	}
}
