package zabbix_test

import (
	"encoding/json"
	zapi "github.com/claranet/go-zabbix-api"
	"log"
	"testing"
)

func TestPassiveProxy(t *testing.T) {
	proxyInterfaces := &zapi.ProxyInterface{IP: "10.1.1.1", Port: "1234", UseIP: 1}
	passiveProxy := &zapi.Proxy{Name: "TestZabbixProxy", Status: zapi.PassiveProxy, Interface: proxyInterfaces}
	log.Printf("%+v", passiveProxy)
	//testCRUDAPIObjectOperations(t, passiveProxy)
	testCreateAPIObject(t, passiveProxy)
	defer testDeleteAPIObject(t, passiveProxy)
	passiveProxy = &zapi.Proxy{ProxyID: passiveProxy.ProxyID}
	testReadAPIObject(t, passiveProxy)
	jsonInterface, err := json.Marshal(passiveProxy.Interface)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(jsonInterface, proxyInterfaces)
	if err != nil {
		t.Fatal(err)
	}
	passiveProxy.Interface = proxyInterfaces
	log.Printf("%+v", passiveProxy)
	testUpdateAPIObject(t, passiveProxy)
}

func TestActiveProxy(t *testing.T) {
	activeProxy := &zapi.Proxy{Name: "TestZabbixProxy", Status: zapi.ActiveProxy, Interface: nil}
	log.Printf("%+v", activeProxy)
	//testCRUDAPIObjectOperations(t, activeProxy)
	testCreateAPIObject(t, activeProxy)
	defer testDeleteAPIObject(t, activeProxy)
	activeProxy = &zapi.Proxy{ProxyID: activeProxy.ProxyID}
	testReadAPIObject(t, activeProxy)
	testUpdateAPIObject(t, activeProxy)
}
