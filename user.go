package zabbix

type (
	// Whether to pause escalation during maintenance periods or not.
	// "debug_mode" in https://www.zabbix.com/documentation/4.0/manual/api/reference/user/object#user_group
	UserType int
)

const (
	ZabbixUser       UserType = 0
	ZabbixAdmin      UserType = 1
	ZabbixSuperAdmin UserType = 2
)

// User represent Zabbix user group object
// https://www.zabbix.com/documentation/4.0/manual/api/reference/user/object
type User struct {
	UserID   string   `json:"userid,omitempty"`
	Alias    string   `json:"alias"`
	Username string   `json:"username"` // Renamed field alias → username in user object from 5.4↑
	Name     string   `json:"name,omitempty"`
	Surname  string   `json:"surname,omitempty"`
	Type     UserType `json:"type,string,omitempty"`
	Url      string   `json:"url,omitempty"`
}

// Users is an array of User
type Users []User

// UsersGet Wrapper for user.get
// https://www.zabbix.com/documentation/4.0/manual/api/reference/user/get
func (api *API) UsersGet(params Params) (res Users, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("user.get", params, &res)
	return
}
