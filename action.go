package zabbix

type Condition struct {
	ConditionType int    `json:"conditiontype,string,omitempty"`
	Operator      int    `json:"operator,string,omitempty"`
	Value         string `json:"value,omitempty"`
}

type Filter struct {
	EvalType    int              `json:"evaltype,string"`
	Conditions  Conditions       `json:"conditions,omitempty"`
	EventSource ActionSourceType `json:"eventsource,string,omitempty"`
	EvalFormula string           `json:"eval_formula,omitempty"`
	Formula     string           `json:"formula,omitempty"`
}

type OpMessageGrp struct {
	UsrGrpID string `json:"usrgrpid,omitempty"`
}

type OpMessage struct {
	DefaultMsg  int    `json:"default_msg,string,omitempty"`
	MediaTypeID string `json:"mediatypeid,omitempty"`
	Message     string `json:"message,omitempty"`
	Subject     string `json:"subject,omitempty"`
}

type OpCommandGrp struct {
	Groupid string `json:"groupid,omitempty"`
}

type OpCommand struct {
	ScriptID string `json:"scriptid,omitempty"`
}

type Operation struct {
	OperationType OperationType `json:"operationtype,string"`
	EscPeriod     string        `json:"esc_period,omitempty"`
	EscStepFrom   int           `json:"esc_step_from,string,omitempty"`
	EscStepTo     int           `json:"esc_step_to,string,omitempty"`
	EvalType      int           `json:"evaltype,string,omitempty"`
	OpMessageGrp  OpMessageGrps `json:"opmessage_grp,omitempty"`
	OpMessage     *OpMessage    `json:"opmessage,omitempty"`
	OpCommandGrp  OpCommandGrps `json:"opcommand_grp,omitempty"`
	OpCommand     *OpCommand    `json:"opcommand,omitempty"`
}

type RecoveryOperation struct {
	OperationType OperationType `json:"operationtype,string,omitempty"`
	OpMessage     OpMessage     `json:"opmessage,omitempty"`
}

type UpdateOperation struct {
	OperationType OperationType `json:"operationtype,string,omitempty"`
	OpMessage     OpMessage     `json:"opmessage,omitempty"`
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

	// TriggerEventSource type
	TriggerEventSource ActionSourceType = 0
	// DiscoveryRuleEventSource rule type
	DiscoveryRuleEventSource ActionSourceType = 1
	// ActiveAgentAutoRegistrationEventSource type
	ActiveAgentAutoRegistrationEventSource ActionSourceType = 2
	// InternalEventSource type
	InternalEventSource ActionSourceType = 3
	// ServiceStatusUpdateEventSource type
	ServiceStatusUpdateEventSource ActionSourceType = 4

	// Different Operation type,

	SendMessageOperationType          = 0
	GlobalScriptOperationType         = 1
	AddHostOperationType              = 2
	RemoveHostOperationType           = 3
	AddToHostGroupOperationType       = 4
	RemoveFromHostGroupOperationType  = 5
	LinkToTemplateOperationType       = 6
	UnlinkFromTemplateOperationType   = 7
	EnableHostOperationType           = 8
	DisableHostOperationType          = 9
	SetHostInventoryModeOperationType = 10
)

// Action represent Zabbix Action Group type returned from Zabbix API
type Action struct {
	ActionIDs          string             `json:"actionid,omitempty"`
	Name               string             `json:"name,omitempty"`
	EventSource        ActionSourceType   `json:"eventsource,string,omitempty"`
	Status             int                `json:"status,string,omitempty"`
	EscPeriod          string             `json:"esc_period,omitempty"`
	Filter             Filter             `json:"filter,omitempty"`
	Operations         Operations         `json:"operations,omitempty"`
	RecoveryOperations RecoveryOperations `json:"recovery_operations,omitempty"`
	UpdateOperations   UpdateOperations   `json:"update_operations,omitempty"`
	PauseSuppressed    int                `json:"pause_suppressed,string,omitempty"`
	NotifyIfCanceled   int                `json:"notify_if_canceled,string,omitempty"`
}

// Actions  is an Array of Action structs.
type Actions []Action

// ActionGet Wrapper for action.get
func (api *API) ActionGet(params Params) (res Actions, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
		params["selectOperations"] = "extend"
		params["selectRecoveryOperations"] = "extend"
		params["selectUpdateOperations"] = "extend"
		params["selectFilter"] = "extend"
		params["filter"] = Filter{EventSource: InternalEventSource}
	}
	err = api.CallWithErrorParse("action.get", params, &res)
	return
}

// ActionCreate Wrapper for action.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/Action/create
func (api *API) ActionCreate(Actions Actions) (err error) {
	response, err := api.CallWithError("action.create", Actions)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	actionids := result["actionids"].([]interface{})

	for i, id := range actionids {
		Actions[i].ActionIDs = id.(string)
	}

	return
}

// ActionUpdate Wrapper for action.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/Action/update
func (api *API) ActionUpdate(Actions Actions) (err error) {
	_, err = api.CallWithError("action.update", Actions)
	return
}

// ActionDelete Wrapper for action.delete
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

// ActionDeleteByIds Wrapper for action.delete
// Use Action's id to delete the Action
// https://www.zabbix.com/documentation/3.2/manual/api/reference/Action/delete
func (api *API) ActionDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("action.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	idss := result["actionids"].([]interface{})
	if len(ids) != len(idss) {
		err = &ExpectedMore{len(ids), len(idss)}
	}

	return
}
