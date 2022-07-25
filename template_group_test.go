package zabbix_test

import (
	"fmt"
	"testing"
	"time"

	zapi "github.com/claranet/go-zabbix-api"
)

func testCreateTemplateGroup(t *testing.T) *zapi.TemplateGroup {
	now := time.Now()
	templateGroups := zapi.TemplateGroups{zapi.TemplateGroup{
		Name: "template group name" + fmt.Sprint(now.UnixNano()),
	}}
	err := testGetAPI(t).TemplateGroupCreate(templateGroups)
	if err != nil {
		t.Fatal(err)
	}
	return &templateGroups[0]
}
func testDeleteTemplateGroup(template *zapi.TemplateGroup, t *testing.T) {
	err := testGetAPI(t).TemplateGroupDelete(zapi.TemplateGroups{*template})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTemplatesGroup(t *testing.T) {
	api := testGetAPI(t)

	template := testCreateTemplateGroup(t)
	if template.GroupIDs == "" {
		t.Fatalf("Template id is empty %#v", template)
	}

	templates, err := api.TemplateGroupGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(templates) == 0 {
		t.Fatal("No templates group were obtained")
	}

	template.Name = "new template group name"
	template.GroupID = template.GroupIDs
	template.GroupIDs = ""
	err = api.TemplateGroupUpdate(zapi.TemplateGroups{*template})
	if err != nil {
		t.Fatal(err)
	}

	testDeleteTemplateGroup(template, t)
}
