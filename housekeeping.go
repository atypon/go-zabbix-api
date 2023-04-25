package zabbix

type HousekeepingSettings struct {
	EventsMode                 int    `json:"hk_events_mode,string"`
	EventsTriggerStoragePeriod string `json:"hk_events_trigger"`
	//EventsService              string `json:"hk_events_service"`
	EventsDataStoragePeriod   string `json:"hk_events_internal"`
	EventsDiscoveryPeriod     string `json:"hk_events_discovery"`
	EventsAutoregPeriod       string `json:"hk_events_autoreg"`
	ServicesMode              int    `json:"hk_services_mode,string"`
	ServicesDataStoragePeriod string `json:"hk_services"`
	AuditMode                 int    `json:"hk_audit_mode,string"`
	AuditStoragePeriod        string `json:"hk_audit"`
	SessionsMode              int    `json:"hk_sessions_mode,string"`
	SessionsStoragePeriod     string `json:"hk_sessions"`
	HistoryMode               int    `json:"hk_history_mode,string"`
	HistoryGlobal             int    `json:"hk_history_global,string"`
	HistoryStoragePeriod      string `json:"hk_history"`
	TrendsMode                int    `json:"hk_trends_mode,string"`
	TrendsGlobal              int    `json:"hk_trends_global,string"`
	TrendsStoragePeriod       string `json:"hk_trends"`
	DBExtension               string `json:"db_extension,omitempty"`
	CompressionStatus         int    `json:"compression_status,string"`
	CompressOlderThan         string `json:"compress_older"`
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
