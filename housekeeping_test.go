package zabbix_test

import (
	"encoding/json"
	"testing"
)

func TestHousekeeping(t *testing.T) {
	houseKeeping, err := _api.HousekeepingGet()
	if err != nil {
		t.Fatal(err)
	}
	jsonObject, err := json.Marshal(houseKeeping)
	t.Logf("housekeeping get result %s", jsonObject)
	err = _api.HousekeepingSet(houseKeeping)
	if err != nil {
		t.Fatal(err)
	}
}
