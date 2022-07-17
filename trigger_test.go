package zabbix_test

import (
	"fmt"
	"testing"

	zapi "github.com/claranet/go-zabbix-api"
)

func testCreateTrigger(item *zapi.Item, host *zapi.Host, t *testing.T) *zapi.Trigger {
	expression := fmt.Sprintf("{%s:%s.last()}=0", host.Host, item.Key)
	triggers := zapi.Triggers{{
		Description: "trigger description",
		Expression:  expression,
	}}
	err := testGetAPI(t).TriggersCreate(triggers)
	if err != nil {
		t.Fatal(err)
	}
	return &triggers[0]
}

func testDeleteTrigger(trigger *zapi.Trigger, t *testing.T) {
	err := testGetAPI(t).TriggersDelete(zapi.Triggers{*trigger})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTrigger(t *testing.T) {
	api := testGetAPI(t)

	group := testCreateHostGroup(t)
	defer testDeleteHostGroup(group, t)

	host := testCreateHost(group, t)
	defer testDeleteHost(host, t)

	app := testCreateApplication(host, t)
	defer testDeleteApplication(app, t)

	item := testCreateItem(app, t)
	defer testDeleteItem(item, t)

	triggerParam := zapi.Params{"hostids": host.HostID}
	res, err := api.TriggersGet(triggerParam)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 0 {
		t.Fatal("Found items")
	}

	trigger := testCreateTrigger(item, host, t)

	trigger.Description = "new trigger name"
	err = api.TriggersUpdate(zapi.Triggers{*trigger})
	if err != nil {
		t.Error(err)
	}

	testDeleteTrigger(trigger, t)
}
