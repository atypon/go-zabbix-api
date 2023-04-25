package zabbix_test

import (
	"testing"

	zapi "github.com/claranet/go-zabbix-api"
)

func testCreateItem(t *testing.T) *zapi.Item { //app *zapi.Application
	items := zapi.Items{{
		//HostID:         app.HostID,
		Key:  "key.lala.laa",
		Name: "name for key",
		// an empty templated created for this test only
		InterfaceID: "30084",
		// the server ID
		HostID: "10528",
		Type:   zapi.ZabbixTrapper,
		Delay:  "0",
		//ApplicationIds: []string{app.ApplicationID},
	}}
	err := testGetAPI(t).ItemsCreate(items)
	if err != nil {
		t.Fatal(err)
	}
	return &items[0]
}

func testDeleteItem(item *zapi.Item, t *testing.T) {
	err := testGetAPI(t).ItemsDelete(zapi.Items{*item})
	if err != nil {
		t.Fatal(err)
	}
}

func TestItems(t *testing.T) {
	api := testGetAPI(t)

	group := testCreateHostGroup(t)
	defer testDeleteHostGroup(group, t)

	host := testCreateHost(group, t)
	defer testDeleteHost(host, t)

	// Applications are deprecated
	/*
		app := testCreateApplication(host, t)
		defer testDeleteApplication(app, t)

		items, err := api.ItemsGetByApplicationID(app.ApplicationID)
		if err != nil {
			t.Fatal(err)
		}
		if len(items) != 0 {
			t.Fatal("Found items")
		}
	*/

	item := testCreateItem(t)

	_, err := api.ItemGetByID(item.ItemID)
	if err != nil {
		t.Fatal(err)
	}

	item.Name = "another name"
	err = api.ItemsUpdate(zapi.Items{*item})
	if err != nil {
		t.Error(err)
	}

	testDeleteItem(item, t)
}
