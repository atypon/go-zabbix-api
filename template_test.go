package zabbix_test

import (
	"testing"

	zapi "github.com/claranet/go-zabbix-api"
)

func testCreateTemplate(hostGroup *zapi.HostGroup, t *testing.T) *zapi.Template {
	// UglyFix the create template API in 6.2 doesn't support group's name in the groups list
	hostGroupTmp := zapi.HostGroups{{GroupID: hostGroup.GroupID}}[0]

	template := zapi.Templates{zapi.Template{
		Host:   "template name",
		Groups: zapi.HostGroups{hostGroupTmp},
	}}
	err := testGetAPI(t).TemplatesCreate(template)
	if err != nil {
		t.Fatal(err)
	}
	return &template[0]
}

func testDeleteTemplate(template *zapi.Template, t *testing.T) {
	err := testGetAPI(t).TemplatesDelete(zapi.Templates{*template})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTemplates(t *testing.T) {
	api := testGetAPI(t)

	hostGroup := testCreateHostGroup(t)

	defer testDeleteHostGroup(hostGroup, t)

	template := testCreateTemplate(hostGroup, t)
	if template.TemplateID == "" {
		t.Errorf("Template id is empty %#v", template)
	}

	templates, err := api.TemplatesGet(zapi.Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(templates) == 0 {
		t.Fatal("No templates were obtained")
	}

	_, err = api.TemplateGetByID(template.TemplateID)
	if err != nil {
		t.Error(err)
	}

	template.Name = "new template name"
	err = api.TemplatesUpdate(zapi.Templates{*template})
	if err != nil {
		t.Error(err)
	}

	testDeleteTemplate(template, t)
}
