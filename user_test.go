package zabbix_test

import (
	zapi "github.com/claranet/go-zabbix-api"
	"testing"
)

func TestUser(t *testing.T) {
	userGroup := &zapi.UserGroup{Name: "UserTestUserGroup", GuiAccess: 2}
	testCreateAPIObject(t, userGroup)
	defer testDeleteAPIObject(t, userGroup)
	role := &zapi.Role{Name: "UserTestRole", Type: zapi.UserRole}
	testCreateAPIObject(t, role)
	defer testDeleteAPIObject(t, role)
	user := &zapi.User{Username: "TestUserName", Name: "TestUser", RoleID: role.GetID(), Groups: []zapi.UserGroupID{zapi.UserGroupID(userGroup.GetID())}}
	testCRUDAPIObjectOperations(t, user)
}
