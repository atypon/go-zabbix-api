package zabbix

type (
	// Whether to pause escalation during maintenance periods or not.
	// "debug_mode" in https://www.zabbix.com/documentation/4.0/manual/api/reference/usergroup/object#user_group
	DebugModeType int
)

const (
	DebugModeDisabled DebugModeType = 0
	DebugModeEnabled  DebugModeType = 1
)

// UserGroup represent Zabbix user group object
// https://www.zabbix.com/documentation/4.0/manual/api/reference/usergroup/object
type UserGroup struct {
	GroupID             string        `json:"usrgrpid,omitempty" zabbix:"id"`
	Name                string        `json:"name"`
	DebugMode           DebugModeType `json:"debug_mode,string,omitempty"`
	GuiAccess           int           `json:"gui_access,string,omitempty"`
	UserStatus          StatusType    `json:"users_status,string,omitempty"`
	UserIDs             []string      `json:"users,omitempty"`
	HostGroupRights     []Permission  `json:"hostgroup_rights,omitempty"`
	TemplateGroupRights []Permission  `json:"templategroup_rights,omitempty"`
}

type Permission struct {
	ID         string          `json:"id"`
	Permission PermissionLevel `json:"permission"`
}

type PermissionLevel int

const (
	AccessDenied PermissionLevel = iota
	ReadOnly
	ReadWrite
)

// UserGroups is an array of UserGroup
type UserGroups []UserGroup

// UserGroupsGet Wrapper for usergroup.get
// https://www.zabbix.com/documentation/4.0/manual/api/reference/usergroup/get
func (api *API) UserGroupsGet(params Params) (res UserGroups, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("usergroup.get", params, &res)
	return
}

func (u *UserGroup) GetID() string {
	return u.GroupID
}

func (u *UserGroup) SetID(ID string) {
	u.GroupID = ID
}

func (u *UserGroup) GetAPIModule() string {
	return "usergroup"
}
