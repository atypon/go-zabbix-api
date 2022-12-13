package zabbix

import (
	"fmt"
)

type Role struct {
	RoleID   string   `json:"roleid,omitempty"`
	Name     string   `json:"name"`
	Type     RoleType `json:"type,string"`
	ReadOnly int      `json:"readonly,string,omitempty"`
}

type RoleType int

const (
	UserRole RoleType = iota + 1
	AdminRole
	SuperAdminRole
)

type Roles []Role

func (api *API) RolesGet(params Params) (res Roles, err error) {
	err = api.CallWithErrorParse("role.get", params, &res)
	return
}

func (api *API) RoleGetByID(id string) (res Role, err error) {
	var roles Roles
	err = api.CallWithErrorParse("role.get", Params{"roleids": id}, &roles)
	if len(roles) == 0 {
		err = fmt.Errorf("role with ID: %s not found", id)
		return
	}
	if err != nil {
		return
	}
	res = roles[0]
	return
}

func (api *API) RoleGetByName(roleName string) (role Role, err error) {
	params := Params{
		"filter": map[string]interface{}{
			"name": roleName,
		},
	}
	roles, err := api.RolesGet(params)
	role = roles[0]
	return
}

func (api *API) RolesCreateAndSetIDs(roles Roles) (err error) {
	response, err := api.CallWithError("role.create", roles)
	if err != nil {
		return
	}
	result := response.Result.(map[string]interface{})
	ids := result["roleids"].([]interface{})
	if len(ids) == 0 {
		return fmt.Errorf("could not create role %s", roles)
	}
	for i, id := range ids {
		roles[i].RoleID = id.(string)
	}
	return
}

func (api *API) RolesDeleteByIDs(ids []string) (err error) {
	_, err = api.CallWithError("role.delete", ids)
	return
}

func (api *API) RolesDeleteByID(id string) (err error) {
	_, err = api.CallWithError("role.delete", []string{id})
	return
}

func (api *API) RolesUpdate(roles Roles) (err error) {
	_, err = api.CallWithError("role.update", roles)
	return
}

func (role Role) String() string {
	return role.Name
}
