package zabbix_test

import (
	"testing"

	zapi "github.com/claranet/go-zabbix-api"
)

func testCreateTemplate(templateGroup *zapi.TemplateGroup, t *testing.T) *zapi.Template {

	templateGroup.Name = ""
	templateGroup.GroupID = templateGroup.GroupIDs
	templateGroup.GroupIDs = ""
	template := zapi.Templates{zapi.Template{
		Host:   "template name",
		Groups: zapi.TemplateGroups{*templateGroup},
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

	templateGroup := testCreateTemplateGroup(t)

	defer testDeleteTemplateGroup(templateGroup, t)

	template := testCreateTemplate(templateGroup, t)
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
