package zabbix

import (
	"strconv"
)

type (
	// Whether to pause escalation during maintenance periods or not.
	// "pause_suppressed" in https://www.zabbix.com/documentation/4.0/manual/api/reference/action/object#action
	PauseType int

	// Filter condition evaluation method.
	// "evaltype" in https://www.zabbix.com/documentation/4.0/manual/api/reference/action/object#action_filter
	ActionEvaluationType int

	// Type of condition.
	// "conditiontype" in https://www.zabbix.com/documentation/4.0/manual/api/reference/action/object#action_filter_condition
	// "conditiontype" in https://www.zabbix.com/documentation/4.0/manual/api/reference/action/object#action_operation_condition
	ActionConditionType int

	// Condition operator.
	// "operator" in https://www.zabbix.com/documentation/4.0/manual/api/reference/action/object#action_filter_condition
	ActionFilterConditionOperator int

	// Type of operation.
	// "operationtype" in https://www.zabbix.com/documentation/4.0/manual/api/reference/action/object#action_operation
	ActionOperationType int

	// Type of operation command.
	// "type" in https://www.zabbix.com/documentation/4.0/manual/api/reference/action/object#action_operation_command
	ActionOperationCommandType int

	// Authentication method used for SSH commands.
	// "authtype" in https://www.zabbix.com/documentation/4.0/manual/api/reference/action/object#action_operation_command
	ActionOperationCommandAuthType int

	// Target on which the custom script operation command will be executed.
	// "execute_on" in https://www.zabbix.com/documentation/4.0/manual/api/reference/action/object#action_operation_command
	ActionOperationCommandExecutorType int
)

const (
	DontPause PauseType = 0
	Pause     PauseType = 1
)

const (
	AndOr  ActionEvaluationType = 0
	And    ActionEvaluationType = 1
	Or     ActionEvaluationType = 2
	Custom ActionEvaluationType = 3
)

const (
	HostGroupCondition                ActionConditionType = 0
	HostCondition                     ActionConditionType = 1
	TriggerCondition                  ActionConditionType = 2
	TriggerNameCondition              ActionConditionType = 3
	TriggerSeverityCondition          ActionConditionType = 4
	TimePeriodCondition               ActionConditionType = 6
	HostIpCondition                   ActionConditionType = 7
	DiscoveredServiceTypeCondition    ActionConditionType = 8
	DiscoveredServicePortCondition    ActionConditionType = 9
	DiscoveryStatusCondition          ActionConditionType = 10
	UptimeOrDowntimeDurationCondition ActionConditionType = 11
	ReceivedValueCondition            ActionConditionType = 12
	HostTemplateCondition             ActionConditionType = 13
	EventAcknowledgedCondition        ActionConditionType = 14
	ApplicationCondition              ActionConditionType = 15
	ProblemIsSuppressedCondition      ActionConditionType = 16
	DiscoveryRuleCondition            ActionConditionType = 18
	DiscoveryCheckCondition           ActionConditionType = 19
	ProxyCondition                    ActionConditionType = 20
	DiscoveryObjectCondition          ActionConditionType = 21
	HostNameCondition                 ActionConditionType = 22
	EventTypeCondition                ActionConditionType = 23
	HostMetadataCondition             ActionConditionType = 24
	EventTagCondition                 ActionConditionType = 25
	EventTagValueCondition            ActionConditionType = 26
)

const (
	Equals                ActionFilterConditionOperator = 0
	DoesNotEqual          ActionFilterConditionOperator = 1
	Contains              ActionFilterConditionOperator = 2
	DoesNotContains       ActionFilterConditionOperator = 3
	In                    ActionFilterConditionOperator = 4
	IsGreaterThanOrEquals ActionFilterConditionOperator = 5
	IsLessThanOrEquals    ActionFilterConditionOperator = 6
	NotIn                 ActionFilterConditionOperator = 7
	Matches               ActionFilterConditionOperator = 8
	DoesNotMatch          ActionFilterConditionOperator = 9
	Yes                   ActionFilterConditionOperator = 10
	No                    ActionFilterConditionOperator = 11
)

const (
	SendMessage               ActionOperationType = 0
	RemoteCommand             ActionOperationType = 1
	AddHost                   ActionOperationType = 2
	RemoveHost                ActionOperationType = 3
	AddToHostGroup            ActionOperationType = 4
	RemoveFromHostGroup       ActionOperationType = 5
	LinkToTemplate            ActionOperationType = 6
	UnlinkFromTemplate        ActionOperationType = 7
	EnableHost                ActionOperationType = 8
	DisableHost               ActionOperationType = 9
	SetHostInventoryMode      ActionOperationType = 10
	NotifyRecoveryAllInvolved ActionOperationType = 11
	NotifyUpdateAllInvolved   ActionOperationType = 12
)

const (
	CustomScript  ActionOperationCommandType = 0
	IpmiCommand   ActionOperationCommandType = 1
	SshCommand    ActionOperationCommandType = 2
	TelnetCommand ActionOperationCommandType = 3
	GlobalScript  ActionOperationCommandType = 4
)

const (
	Password  ActionOperationCommandAuthType = 0
	PublicKey ActionOperationCommandAuthType = 1
)

const (
	AgentExecutor  ActionOperationCommandExecutorType = 0
	ServerExecutor ActionOperationCommandExecutorType = 1
	ProxyExecutor  ActionOperationCommandExecutorType = 2
)

// Action represent Zabbix Action type returned from Zabbix API
// https://www.zabbix.com/documentation/4.0/manual/api/reference/action/object
// https://www.zabbix.com/documentation/5.0/manual/api/reference/action/object
type Action struct {
	ActionID string `json:"actionid,omitempty"`
	//Period      string    `json:"esc_period"`
	EventSource EventType `json:"eventsource,omitempty"` // NOTE: Can not update
	Name        string    `json:"name"`
	//DefaultMessage  string     `json:"def_longdata"`  // NOTE: no support on Zabbix 5.0 onward
	//DefaultSubject  string     `json:"def_shortdata"` // NOTE: no support on Zabbix 5.0 onward
	//RecoveryMessage string     `json:"r_longdata"`    // NOTE: no support on Zabbix 5.0 onward
	//RecoverySubject string     `json:"r_shortdata"`   // NOTE: no support on Zabbix 5.0 onward
	//AckMessage      string     `json:"ack_longdata"`  // NOTE: no support on Zabbix 5.0 onward
	//AckSubject      string     `json:"ack_shortdata"` // NOTE: no support on Zabbix 5.0 onward
	Status          StatusType `json:"status,omitempty,string"`
	PauseSuppressed *PauseType `json:"pause_suppressed,omitempty,string"`

	Filter             ActionFilter             `json:"filter,omitempty"`
	Operations         ActionOperations         `json:"operations,omitempty"`
	RecoveryOperations ActionRecoveryOperations `json:"recoveryOperations,omitempty"`
	UpdateOperations   ActionUpdateOperations   `json:"acknowledgeOperations,omitempty"`
}

type Actions []Action

type ActionFilter struct {
	Conditions     ActionFilterConditions `json:"conditions"`
	EvaluationType ActionEvaluationType   `json:"evaltype,string"`
	Formula        string                 `json:"formula,omitempty"`
}

type ActionFilterCondition struct {
	ConditionID   string              `json:"conditionid,omitempty"`
	ConditionType ActionConditionType `json:"conditiontype,string"`
	Value         string              `json:"value"`
	//Value2        string                        `json:"value2,omitempty"`
	FormulaID string                        `json:"formulaid,omitempty"`
	Operator  ActionFilterConditionOperator `json:"operator,string"`
}

type ActionFilterConditions []ActionFilterCondition

type ActionOperation struct {
	OperationID       string                           `json:"operationid,omitempty"`
	OperationType     ActionOperationType              `json:"operationtype,string"`
	ActionID          string                           `json:"actionid,omitempty"`
	EvaluationType    ActionEvaluationType             `json:"evaltype,omitempty,string"`
	Command           *ActionOperationCommand          `json:"opcommand,omitempty"`
	CommandHostGroups ActionOperationCommandHostGroups `json:"opcommand_grp,omitempty"`
	CommandHosts      ActionOperationCommandHosts      `json:"opcommand_hst,omitempty"`
	Conditions        ActionOperationConditions        `json:"opconditions,omitempty"`
	HostGroups        ActionOperationHostGroups        `json:"opgroup,omitempty"`
	Message           *ActionOperationMessage          `json:"opmessage,omitempty"`
	MessageUserGroups ActionOperationMessageUserGroups `json:"opmessage_grp,omitempty"`
	MessageUsers      ActionOperationMessageUsers      `json:"opmessage_usr,omitempty"`
	Templates         ActionOperationTemplates         `json:"optemplate,omitempty"`
	Inventory         *ActionOperationInventory        `json:"opinventory,omitempty"`
}

type ActionOperations []ActionOperation

type ActionOperationCommand struct {
	OperationID string                              `json:"operationid,omitempty"`
	Type        ActionOperationCommandType          `json:"type,string"`
	Command     string                              `json:"command,omitempty"`
	AuthType    *ActionOperationCommandAuthType     `json:"authtype,omitempty,string"`
	ExecuteOn   *ActionOperationCommandExecutorType `json:"execute_on,omitempty,string"`
	Username    string                              `json:"username,omitempty"`
	Password    string                              `json:"password,omitempty"`
	Port        string                              `json:"port,omitempty"`
	PrivateKey  string                              `json:"privatekey,omitempty"`
	PublicKey   string                              `json:"publickey,omitempty"`
	ScriptID    string                              `json:"scriptid,omitempty"`
}

type ActionOperationCommandHostGroup struct {
	CommandHostGroupID string `json:"opcommand_grpid,omitempty"`
	OperationID        string `json:"operationid,omitempty"`
	GroupID            string `json:"groupid"`
}

type ActionOperationCommandHostGroups []ActionOperationCommandHostGroup

type ActionOperationCommandHost struct {
	CommandHostID string `json:"opcommand_hstid,omitempty"`
	OperationID   string `json:"operationid,omitempty"`
	HostID        string `json:"hostid"`
}

type ActionOperationCommandHosts []ActionOperationCommandHost

type ActionOperationCondition struct {
	OperationID string                        `json:"operationid,omitempty"`
	ConditionID string                        `json:"opconditionid,omitempty"`
	Condition   ActionConditionType           `json:"conditiontype,string"`
	Value       string                        `json:"value"`
	Operator    ActionFilterConditionOperator `json:"operator,string"`
}

type ActionOperationConditions []ActionOperationCondition

type ActionOperationHostGroup struct {
	OperationID string `json:"operationid,omitempty"`
	GroupID     string `json:"groupid"`
}

type ActionOperationHostGroups []ActionOperationHostGroup

type ActionOperationMessage struct {
	OperationID    string `json:"operationid,omitempty"`
	DefaultMessage string `json:"default_msg"`
	MediaTypeID    string `json:"mediatypeid"`
	Message        string `json:"message"`
	Subject        string `json:"subject"`
}

type ActionOperationMessageUserGroup struct {
	OperationID string `json:"operationid,omitempty"`
	UserGroupID string `json:"usrgrpid"`
}

type ActionOperationMessageUserGroups []ActionOperationMessageUserGroup

type ActionOperationMessageUser struct {
	OperationID string `json:"operationid,omitempty"`
	UserID      string `json:"userid"`
}

type ActionOperationMessageUsers []ActionOperationMessageUser

type ActionOperationTemplate struct {
	OperationID string `json:"operationid,omitempty"`
	TemplateID  string `json:"templateid"`
}

type ActionOperationTemplates []ActionOperationTemplate

type ActionOperationInventory struct {
	OperationID   string `json:"operationid,omitempty"`
	InventoryMode string `json:"inventory_mode"`
}

type ActionRecoveryOperation struct {
	OperationID       string                           `json:"operationid,omitempty"`
	OperationType     ActionOperationType              `json:"operationtype,string"`
	ActionID          string                           `json:"actionid,omitempty"`
	Command           *ActionOperationCommand          `json:"opcommand,omitempty"`
	CommandHostGroups ActionOperationCommandHostGroups `json:"opcommand_grp,omitempty"`
	CommandHosts      ActionOperationCommandHosts      `json:"opcommand_hst,omitempty"`
	Message           *ActionOperationMessage          `json:"opmessage,omitempty"`
	MessageUserGroups ActionOperationMessageUserGroups `json:"opmessage_grp,omitempty"`
	MessageUsers      ActionOperationMessageUsers      `json:"opmessage_usr,omitempty"`
}

type ActionRecoveryOperations []ActionRecoveryOperation

type ActionUpdateOperation struct {
	OperationID       string                           `json:"operationid,omitempty"`
	OperationType     ActionOperationType              `json:"operationtype,string"`
	Command           *ActionOperationCommand          `json:"opcommand,omitempty"`
	CommandHostGroups ActionOperationCommandHostGroups `json:"opcommand_grp,omitempty"`
	CommandHosts      ActionOperationCommandHosts      `json:"opcommand_hst,omitempty"`
	Message           *ActionOperationMessage          `json:"opmessage,omitempty"`
	MessageUserGroups ActionOperationMessageUserGroups `json:"opmessage_grp,omitempty"`
	MessageUsers      ActionOperationMessageUsers      `json:"opmessage_usr,omitempty"`
}

type ActionUpdateOperations []ActionUpdateOperation

// ActionsGet Wrapper for action.get
// https://www.zabbix.com/documentation/4.0/manual/api/reference/action/get
func (api *API) ActionsGet(params Params) (res Actions, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("action.get", params, &res)
	return
}

// ActionGetByID Gets action by Id only if there is exactly 1 matching action.
func (api *API) ActionGetByID(id string) (res *Action, err error) {
	params := Params{
		"actionids":                   id,
		"selectFilter":                "extend",
		"selectOperations":            "extend",
		"selectRecoveryOperations":    "extend",
		"selectAcknowledgeOperations": "extend",
	}
	actions, err := api.ActionsGet(params)
	if err != nil {
		return
	}

	if len(actions) == 1 {
		res = &actions[0]
	} else {
		e := ExpectedOneResult(len(actions))
		err = &e
	}
	return
}

// ActionsCreate Wrapper for action.create
// https://www.zabbix.com/documentation/4.0/manual/api/reference/action/create
func (api *API) ActionsCreate(actions Actions) (err error) {
	response, err := api.CallWithError("action.create", actions)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	actionids := result["actionids"].([]interface{})
	for i, id := range actionids {
		switch id.(type) {
		case string:
			actions[i].ActionID = id.(string)
		case float64:
			actions[i].ActionID = strconv.Itoa(int(id.(float64)))
		}
	}
	return
}

// ActionsUpdate Wrapper for action.update
// https://www.zabbix.com/documentation/4.0/manual/api/reference/action/update
func (api *API) ActionsUpdate(actions Actions) (err error) {
	_, err = api.CallWithError("action.update", actions)
	return
}

// ActionsDelete Wrapper for action.delete
// Cleans ActionId in all actions elements if call succeed.
// https://www.zabbix.com/documentation/4.0/manual/api/reference/action/delete
func (api *API) ActionsDelete(actions Actions) (err error) {
	ids := make([]string, len(actions))
	for i, action := range actions {
		ids[i] = action.ActionID
	}

	err = api.ActionsDeleteByIds(ids)
	if err == nil {
		for i := range actions {
			actions[i].ActionID = ""
		}
	}
	return
}

// ActionsDeleteByIds Wrapper for action.delete
// https://www.zabbix.com/documentation/4.0/manual/api/reference/action/delete
func (api *API) ActionsDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("action.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	actionids := result["actionids"].([]interface{})
	if len(ids) != len(actionids) {
		err = &ExpectedMore{len(ids), len(actionids)}
	}
	return
}
