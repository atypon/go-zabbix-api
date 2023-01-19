package zabbix_test

import (
	"github.com/claranet/go-zabbix-api"
	"testing"
)

func TestHostInterface_(t *testing.T) {
	group := testCreateHostGroup(t)
	defer testDeleteHostGroup(group, t)
	host := testCreateHost(group, t)
	defer testDeleteHost(host, t)
	snmpDetails := &zabbix.SNMPDetails{
		Version:   1,
		Community: "test",
	}
	hostInterface := &zabbix.HostInterface{
		HostID:  host.HostID,
		DNS:     "",
		IP:      "10.1.1.1",
		UseIP:   1,
		Port:    "1234",
		Type:    zabbix.JMXInterface,
		Main:    1,
		Details: snmpDetails,
	}
	testCRUDAPIObjectOperations(t, hostInterface)
}
