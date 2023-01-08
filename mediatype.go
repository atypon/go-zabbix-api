package zabbix

type MediaType struct {
	MediaID           string         `json:"mediatypeid,omitempty" zabbix:"id"`
	MediaName         string         `json:"name"`
	MediaKind         MediaKind      `json:"type,string"`
	Disabled          int            `json:"status,string"`
	Description       string         `json:"description,omitempty"`
	ScriptExecPath    string         `json:"exec_path,omitempty"`
	ScriptParams      string         `json:"exec_params,omitempty"`
	SMTPAuthPassword  string         `json:"passwd,omitempty"`
	SMTPAuthUser      string         `json:"username,omitempty"`
	SMTPFromEmail     string         `json:"smtp_email,omitempty"`
	SMTPHelo          string         `json:"smtp_helo,omitempty"`
	SMTPServer        string         `json:"smtp_server,omitempty"`
	SMTPPort          string         `json:"smtp_port,omitempty"`
	SMTPSecurity      string         `json:"smtp_security,omitempty"`
	WebhookScript     string         `json:"script,omitempty"`
	WebhookTimeout    string         `json:"timeout,omitempty"`
	WebhookParameters []WebhookParam `json:"parameters,omitempty"`
}

type WebhookParam struct {
	Name  string
	Value string
}
type MediaKind int

const (
	EmailMedia MediaKind = iota
	ScriptMedia
	SMSMedia
	WebhookMedia = iota + 1
)

// Medias is an array of Media
type Medias []MediaType

func (api *API) MediaGet(params Params) (res Medias, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("mediatype.get", params, &res)
	return
}

func (mediaType *MediaType) GetID() string {
	return mediaType.MediaID
}

func (mediaType *MediaType) SetID(id string) {
	mediaType.MediaID = id
}

func (mediaType *MediaType) GetAPIModule() string {
	return "mediatype"
}
