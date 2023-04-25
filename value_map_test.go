package zabbix_test

import (
	zapi "github.com/claranet/go-zabbix-api"
	"testing"
)

func TestValueMap(t *testing.T) {
	group := testCreateHostGroup(t)
	defer testDeleteHostGroup(group, t)
	host := testCreateHost(group, t)
	defer testDeleteHost(host, t)
	mappings := []zapi.ValueMapMapping{{Value: "x", Newvalue: "y", Type: zapi.ExactMatchMapping}}
	valueMap := &zapi.ValueMap{Name: "TestValueMapping", HostID: host.HostID, Mappings: mappings}
	//testCRUDAPIObjectOperations(t, valueMap)
	testCreateAPIObject(t, valueMap)
	defer testDeleteAPIObject(t, valueMap)
	testReadAPIObject(t, valueMap)
	// For some reason, API doesn't expect hostid to be set for the update method
	valueMap.HostID = ""
	testUpdateAPIObject(t, valueMap)
}
