package zabbix

type Condition struct {
	ConditionType int    `json:"conditiontype"`
	Operator      int    `json:"operator"`
	Value         string `json:"value"`
}

type Filter struct {
	EvalType   int        `json:"evaltype"`
	Conditions Conditions `json:"conditions"`
}

type OpMessageGrp struct {
	UsrGrpID string `json:"usrgrpid"`
}

type OpMessage struct {
	DefaultMsg  int    `json:"default_msg"`
	MediaTypeID string `json:"mediatypeid,omitempty"`
	Message     string `json:"message,omitempty"`
	Subject     string `json:"subject,omitempty"`
}

type OpCommandGrp struct {
	Groupid string `json:"groupid"`
}

type OpCommand struct {
	ScriptID string `json:"scriptid"`
}

type Operation struct {
	OperationType OperationType `json:"operationtype,omitempty"`
	EscPeriod     string        `json:"esc_period,omitempty"`
	EscStepFrom   int           `json:"esc_step_from,omitempty"`
	EscStepTo     int           `json:"esc_step_to,omitempty"`
	EvalType      int           `json:"evaltype,omitempty"`
	OpMessageGrp  OpMessageGrps `json:"opmessage_grp,omitempty"`
	OpMessage     OpMessage     `json:"opmessage,omitempty"`
	OpCommandGrp  OpCommandGrps `json:"opcommand_grp,omitempty"`
	OpCommand     OpCommand     `json:"opcommand,omitempty"`
}

type RecoveryOperation struct {
	OperationType string    `json:"operationtype"`
	OpMessage     OpMessage `json:"opmessage"`
}

type UpdateOperation struct {
	OperationType string    `json:"operationtype"`
	OpMessage     OpMessage `json:"opmessage,omitempty"`
}

type Filters []Filter
type OpMessageGrps []OpMessageGrp
type OpCommandGrps []OpCommandGrp
type Conditions []Condition
type Operations []Operation
type UpdateOperations []UpdateOperation
type RecoveryOperations []RecoveryOperation

type (
	// ActionSourceType type of the action source
	ActionSourceType int
	OperationType    int
)

const (
	// Different action source type, see :
	// - "source" in https://www.zabbix.com/documentation/6.2/manual/api/reference/item/object

	// TriggerSource type
	TriggerSource ActionSourceType = 0
	// DiscoveryRuleSource rule type
	DiscoveryRuleSource ActionSourceType = 1
	// ActiveAgentAutoRegistrationSource type
	ActiveAgentAutoRegistrationSource ActionSourceType = 2
	// InternalSource type
	InternalSource ActionSourceType = 3
	// ServiceStatusUpdate type
	ServiceStatusUpdate ActionSourceType = 4

	// Different Operation type,

	SendMessageOperation          = 0
	GlobalScriptOperation         = 1
	AddHostOperation              = 2
	RemoveHostOperation           = 3
	AddToHostGroupOperation       = 4
	RemoveFromHostGroupOperation  = 5
	LinkToTemplateOperation       = 6
	UnlinkFromTemplateOperation   = 7
	EnableHostOperation           = 8
	DisableHostOperation          = 9
	SetHostInventoryModeOperation = 10
)

// Action represent Zabbix Action Group type returned from Zabbix API
type Action struct {
	ActionIDs          string             `json:"actionids,omitempty"`
	Name               string             `json:"name,omitempty"`
	EventSource        ActionSourceType   `json:"eventsource,omitempty"`
	Status             int                `json:"status,omitempty"`
	EscPeriod          string             `json:"esc_period,omitempty"`
	Filter             Filters            `json:"filter,omitempty"`
	Operations         Operations         `json:"operations,omitempty"`
	RecoveryOperations RecoveryOperations `json:"recovery_operations,omitempty"`
	UpdateOperations   UpdateOperations   `json:"update_operations,omitempty"`
	PauseSuppressed    string             `json:"pause_suppressed,omitempty"`
	NotifyIfCanceled   string             `json:"notify_if_canceled,omitempty"`
}

// Actions  is an Array of Action structs.
type Actions []Action

// ActionGet Wrapper for Action.get
func (api *API) ActionGet(params Params) (res Actions, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
		params["selectOperations"] = "extend"
		params["selectRecoveryOperations"] = "extend"
		params["selectUpdateOperations"] = "extend"
		params["selectFilter"] = "extend"
	}
	err = api.CallWithErrorParse("Action.get", params, &res)
	return
}

// ActionCreate Wrapper for Action.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/Action/create
func (api *API) ActionCreate(Actions Actions) (err error) {
	_, err = api.CallWithError("Action.create", Actions)
	if err != nil {
		return
	}
	return
}

// ActionUpdate Wrapper for Action.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/Action/update
func (api *API) ActionUpdate(Actions Actions) (err error) {
	_, err = api.CallWithError("Action.update", Actions)
	return
}

// ActionDelete Wrapper for Action.delete
// Cleans ApplicationID in all apps elements if call succeed.
// https://www.zabbix.com/documentation/3.2/manual/api/reference/Action/delete
func (api *API) ActionDelete(Actions Actions) (err error) {
	ActionIds := make([]string, len(Actions))
	for i, Action := range Actions {
		ActionIds[i] = Action.ActionIDs
	}

	err = api.ActionDeleteByIds(ActionIds)
	if err == nil {
		for i := range Actions {
			Actions[i].ActionIDs = ""
		}
	}
	return
}

// ActionDeleteByIds Wrapper for Action.delete
// Use Action's id to delete the Action
// https://www.zabbix.com/documentation/3.2/manual/api/reference/Action/delete
func (api *API) ActionDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("Action.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	idss := result["groupids"].([]interface{})
	if len(ids) != len(idss) {
		err = &ExpectedMore{len(ids), len(idss)}
	}

	return
}
