package zabbix_test

import (
	zapi "github.com/claranet/go-zabbix-api"
	"testing"
)

func TestRole(t *testing.T) {
	api := testGetAPI(t)

	testRoleName := "TestRole"

	role := zapi.Role{Name: testRoleName, Type: zapi.UserRole}
	roles := zapi.Roles{role}

	err := api.RolesCreateAndSetIDs(roles)
	if err != nil {
		t.Fatal(err)
	}
	id := roles[0].RoleID
	role, err = api.RoleGetByID(id)
	if err != nil {
		t.Fatal(err)
	}

	role.Type = zapi.AdminRole
	err = api.RolesUpdate(zapi.Roles{role})
	if err != nil {
		t.Fatal(err)
	}

	role, err = api.RoleGetByName(testRoleName)
	if err != nil {
		t.Fatal(err)
	}
	if role.Type != zapi.AdminRole {
		t.Fatal("Updating role type failed")
	}

	err = api.RolesDeleteByID(role.RoleID)
	if err != nil {
		t.Fatal(err)
	}
}
