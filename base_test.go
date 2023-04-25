package zabbix_test

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"

	zapi "github.com/claranet/go-zabbix-api"
)

var (
	_host string
	_api  *zapi.API
)

func init() {
	rand.Seed(time.Now().UnixNano())

	var err error
	_host, err = os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	_host += "-testing"

	if os.Getenv("TEST_ZABBIX_URL") == "" {
		log.Fatal("Set environment variables TEST_ZABBIX_URL (and optionally TEST_ZABBIX_USER and TEST_ZABBIX_PASSWORD)")
	}
	_api, err = createAPIClient()
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("DEBUG") != "" {
		_api.Logger = log.Default()
	}
}

func testGetHost() string {
	return _host
}

func testGetAPI(t *testing.T) *zapi.API {
	if _api != nil {
		return _api
	}
	var err error
	_api, err = createAPIClient()
	if err != nil {
		t.Fatal(err)
	}
	return _api
}

func createAPIClient() (api *zapi.API, err error) {
	url, user, password := os.Getenv("TEST_ZABBIX_URL"), os.Getenv("TEST_ZABBIX_USER"), os.Getenv("TEST_ZABBIX_PASSWORD")
	_api = zapi.NewAPI(url)
	_api.SetClient(http.DefaultClient)
	v := os.Getenv("TEST_ZABBIX_VERBOSE")
	if v != "" && v != "0" {
		_api.Logger = log.New(os.Stderr, "[zabbix] ", 0)
	}

	if user != "" {
		auth, err := _api.Login(user, password)
		if err != nil {
			return nil, err
		}
		if auth == "" {
			return nil, fmt.Errorf("login failed")
		}
	}
	return _api, nil
}

func TestBadCalls(t *testing.T) {
	api := testGetAPI(t)
	res, err := api.Call("", nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.Error.Code != -32602 && res.Error.Code != -32600 {
		t.Errorf("Expected code -32602/-32600, got %s", res.Error)
	}
}

func TestVersion(t *testing.T) {
	api := testGetAPI(t)
	v, err := api.Version()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Zabbix version %s", v)
	if !regexp.MustCompile(`^\d\.\d\.\d+$`).MatchString(v) {
		t.Errorf("Unexpected version: %s", v)
	}
}
