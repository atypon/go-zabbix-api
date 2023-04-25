package zabbix

// HostID in value map is the template ID to link to

type ValueMap struct {
	ValueMapID string            `json:"valuemapid,omitempty" zabbix:"id"`
	Name       string            `json:"name"`
	Mappings   []ValueMapMapping `json:"mappings"`
	HostID     string            `json:"hostid,omitempty"`
	UUID       string            `json:"uuid,omitempty"`
}
type ValueMapMapping struct {
	Type     ValueMapMappingType `json:"type,string,omitempty"`
	Value    string              `json:"value"`
	Newvalue string              `json:"newvalue"`
}
type ValueMapMappingType int

const (
	ExactMatchMapping ValueMapMappingType = iota
	GreaterOrEqualMapping
	LessOrEqualMapping
	ValueInRangeMapping
	RegexMatchMapping
	DefaultValueMapping
)

func (api *API) ValueMapGet(params Params) (res Medias, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("valuemap.get", params, &res)
	return
}

func (v *ValueMap) GetID() string {
	return v.ValueMapID
}

func (v *ValueMap) SetID(id string) {
	v.ValueMapID = id
}

func (v *ValueMap) GetAPIModule() string {
	return "valuemap"
}

func (v *ValueMap) GetExtraParams() Params {
	return Params{"selectMappings": "extend"}
}
