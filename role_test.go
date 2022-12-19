package zabbix_test

import (
	"fmt"
	zapi "github.com/claranet/go-zabbix-api"
	"testing"
)

func TestRole(t *testing.T) {
	api := testGetAPI(t)

	testRoleName := "TestRole"

	role := &zapi.Role{Name: testRoleName, Type: zapi.UserRole}
	//roles := zapi.Roles{*role}

	err := api.CreateAPIObject(role)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteRole(t, role)
	fmt.Println(role)
	role = &zapi.Role{RoleID: role.GetID()}
	err = api.ReadAPIObject(role)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(role)

	role.Type = zapi.AdminRole
	//err = api.RolesUpdate(zapi.Roles{*role})
	err = api.UpdateAPIObject(role)
	if err != nil {
		t.Fatal(err)
	}

	err = api.ReadAPIObject(role)
	if err != nil {
		t.Fatal(err)
	}
	if role.Type != zapi.AdminRole {
		t.Fatal("Updating role type failed")
	}
}

func deleteRole(t *testing.T, object zapi.APIObject) {
	api := testGetAPI(t)
	// err := api.RolesDeleteByID(id)
	err := api.DeleteAPIObject(object)
	if err != nil {
		t.Fatal(err)
	}
	err = api.ReadAPIObject(object)
	if err == nil {
		t.Fatal("Could not delete object")
	}

}
