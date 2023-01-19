package zabbix_test

import (
	zapi "github.com/claranet/go-zabbix-api"
	"testing"
)

func TestMacro(t *testing.T) {
	group := testCreateHostGroup(t)
	defer testDeleteHostGroup(group, t)
	host := testCreateHost(group, t)
	defer testDeleteHost(host, t)
	macro := &zapi.Macro{MacroName: "{$TEST_MACRO}", Value: "TestMacroValue", HostID: host.HostID}
	//testCRUDAPIObjectOperations(t, macro)
	testCreateAPIObject(t, macro)
	defer testDeleteAPIObject(t, macro)
	testReadAPIObject(t, macro)
	// can't update host id of a macro
	macro.HostID = ""
	testUpdateAPIObject(t, macro)
}
