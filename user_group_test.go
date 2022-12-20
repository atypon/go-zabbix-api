package zabbix_test

import (
	zapi "github.com/claranet/go-zabbix-api"
	"testing"
)

func TestUserGroup(t *testing.T) {
	group := &zapi.UserGroup{Name: "TestUserGroup"}
	testCRUDAPIObjectOperations(t, group)
}
