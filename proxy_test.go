package zabbix_test

import (
	zapi "github.com/claranet/go-zabbix-api"
	"testing"
)

func TestProxy(t *testing.T) {
	//proxyInterfaces := &zapi.ProxyInterface{IP: "10.1.1.1", Port: "1234", UseIP: 1}
	proxy := &zapi.Proxy{Name: "TestZabbixProxy", Status: zapi.ActiveProxy} //, Interface: proxyInterfaces}
	//testCRUDAPIObjectOperations(t, proxy)
	testCreateAPIObject(t, proxy)
	defer testDeleteAPIObject(t, proxy)
	proxy = &zapi.Proxy{ProxyID: proxy.ProxyID}
	testReadAPIObject(t, proxy)
	testUpdateAPIObject(t, proxy)

}
