package zabbix_test

import (
	zapi "github.com/claranet/go-zabbix-api"
	"testing"
)

func TestMediaType(t *testing.T) {
	mediaType := &zapi.MediaType{MediaName: "TestMediaType", MediaKind: zapi.ScriptMedia, ScriptExecPath: "/script.sh"}
	testCRUDAPIObjectOperations(t, mediaType)
	mediaType = &zapi.MediaType{MediaName: "TestMediaType", MediaKind: zapi.WebhookMedia, WebhookScript: "some js script"}
	testCRUDAPIObjectOperations(t, mediaType)
	mediaType = &zapi.MediaType{
		MediaName: "TestMediaType", MediaKind: zapi.EmailMedia,
		SMTPServer:    "smtp.example.com",
		SMTPHelo:      "example.com",
		SMTPFromEmail: "test@example.com",
	}
	testCRUDAPIObjectOperations(t, mediaType)
}
