package zabbix_test

import (
	"fmt"
	"testing"
	"time"

	zapi "github.com/claranet/go-zabbix-api"
)

func testCreateAction(t *testing.T) *zapi.Action {
	now := time.Now()
	Actions := zapi.Actions{zapi.Action{
		Name:        "action name" + fmt.Sprint(now.UnixNano()),
		EventSource: zapi.InternalSource,
		Operations: zapi.Operations{
			zapi.Operation{
				OperationType: zapi.SendMessageOperation,
				OpMessage: zapi.OpMessage{
					Message:     "hi",
					MediaTypeID: "1",
				},
				OpMessageGrp: zapi.OpMessageGrps{
					zapi.OpMessageGrp{
						UsrGrpID: "7",
					},
				},
			},
			zapi.Operation{
				OperationType: zapi.GlobalScriptOperation,
				OpCommand: zapi.OpCommand{
					ScriptID: "3",
				},
				OpCommandGrp: zapi.OpCommandGrps{
					zapi.OpCommandGrp{
						Groupid: "3",
					},
				},
			},
		},
	}}
	err := testGetAPI(t).ActionCreate(Actions)
	if err != nil {
		t.Fatal(err)
	}
	return &Actions[0]
}
func testDeleteAction(action *zapi.Action, t *testing.T) {
	err := testGetAPI(t).ActionDelete(zapi.Actions{*action})
	if err != nil {
		t.Fatal(err)
	}
}

func TestActionsGroup(t *testing.T) {
	api := testGetAPI(t)

	action := testCreateAction(t)
	if action.ActionIDs == "" {
		t.Fatalf("action id is empty %#v", action)
	}

	actions, err := api.ActionGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(actions) == 0 {
		t.Fatal("No actions group were obtained")
	}

	action.Name = "new action group name"
	err = api.ActionUpdate(zapi.Actions{*action})
	if err != nil {
		t.Fatal(err)
	}

	testDeleteAction(action, t)
}
