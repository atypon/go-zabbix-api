package zabbix_test

import (
	"testing"

	zapi "github.com/claranet/go-zabbix-api"
)

func TestUsersGet(t *testing.T) {
	api := testGetAPI(t)

	params := zapi.Params{
		"filter": map[string]interface{}{
			"alias":    "Admin", // Under 5.4
			"username": "Admin", // 5.4 or higher
		},
	}
	users, err := api.UsersGet(params)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 1 {
		t.Errorf("Bad users: %#v", users)
	}
}
