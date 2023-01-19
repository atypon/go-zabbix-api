package zabbix_test

import (
	"fmt"
	"math/rand"
	"testing"

	zapi "github.com/claranet/go-zabbix-api"
)

func testCreateHost(group *zapi.HostGroup, t *testing.T) *zapi.Host {
	name := fmt.Sprintf("%s-%d", testGetHost(), rand.Int())
	iface := zapi.HostInterface{DNS: name, Port: "42", Type: zapi.AgentInterface, UseIP: 0, Main: 1}
	host := &zapi.Host{
		Host:       name,
		Name:       "Name for " + name,
		GroupIds:   zapi.HostGroupIDs{{group.GroupID}},
		Interfaces: zapi.HostInterfaces{iface},
	}

	err := testGetAPI(t).CreateAPIObject(host)
	if err != nil {
		t.Fatal(err)
	}
	return host
}

func testDeleteHost(host *zapi.Host, t *testing.T) {
	err := testGetAPI(t).DeleteAPIObject(host)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHosts(t *testing.T) {
	//api := testGetAPI(t)

	group := testCreateHostGroup(t)
	defer testDeleteHostGroup(group, t)
	templateGroup := testCreateTemplateGroup(t)
	defer testDeleteTemplateGroup(templateGroup, t)
	template := testCreateTemplate(templateGroup, t)
	defer testDeleteTemplate(template, t)

	host := testCreateHost(group, t)
	defer testDeleteAPIObject(t, host)
	testReadAPIObject(t, host)
	host.TemplateIDs = zapi.TemplateIDs{zapi.TemplateID{TemplateID: template.TemplateID}}
	//testCRUDAPIObjectOperations(t, host)
	testUpdateAPIObject(t, host)
}
