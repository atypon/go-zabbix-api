package zabbix

type HousekeepingSettings struct {
	EventsMode                 int    `json:"hk_events_mode,string,omitempty"`
	EventsTriggerStoragePeriod string `json:"hk_events_trigger,omitempty"`
	//EventsService              string `json:"hk_events_service,omitempty"`
	EventsDataStoragePeriod   string `json:"hk_events_internal,omitempty"`
	EventsDiscoveryPeriod     string `json:"hk_events_discovery,omitempty"`
	EventsAutoregPeriod       string `json:"hk_events_autoreg,omitempty"`
	ServicesMode              int    `json:"hk_services_mode,string,omitempty"`
	ServicesDataStoragePeriod string `json:"hk_services,omitempty"`
	AuditMode                 int    `json:"hk_audit_mode,string,omitempty"`
	AuditStoragePeriod        string `json:"hk_audit,omitempty"`
	SessionsMode              int    `json:"hk_sessions_mode,string,omitempty"`
	SessionsStoragePeriod     string `json:"hk_sessions,omitempty"`
	HistoryMode               int    `json:"hk_history_mode,string,omitempty"`
	HistoryGlobal             int    `json:"hk_history_global,string,omitempty"`
	HistoryStoragePeriod      string `json:"hk_history,omitempty"`
	TrendsMode                int    `json:"hk_trends_mode,string,omitempty"`
	TrendsGlobal              int    `json:"hk_trends_global,string,omitempty"`
	TrendsStoragePeriod       string `json:"hk_trends,omitempty"`
	DBExtension               string `json:"db_extension,omitempty"`
	CompressionStatus         int    `json:"compression_status,string,omitempty"`
	CompressOlderThan         string `json:"compress_older,omitempty"`
	CompressionAvailability   int    `json:"compression_availability,string,omitempty"`
}

func (api *API) HousekeepingGet() (houseKeeping *HousekeepingSettings, err error) {
	houseKeeping = &HousekeepingSettings{}
	err = api.CallWithErrorParse("housekeeping.get", Params{"output": "extend"}, houseKeeping)
	return
}

func (api *API) HousekeepingSet(houseKeeping *HousekeepingSettings) error {
	_, err := api.CallWithError("housekeeping.update", houseKeeping)
	return err
}
