package main

import (
	"github.com/siye1982/eagleye-health/config"
	"github.com/siye1982/eagleye-health/registry"
	"time"
)

func main() {

	config.EtcdHosts = "http://192.168.10.235:2379,http://192.168.10.236:2379"
	config.GroupName = "packet"
	config.HeartbeatConfig = `{"sss":"aabbcc"}`
	config.InitEtcdClient()
	registry.Start()
	time.Sleep(1 * time.Second)

	// 独立线程模拟计数
	go func() {
		for {
			registry.TpmCounter()
			registry.TtCounter()
		}
	}()


	//registry.Start()
	time.Sleep(1000 * time.Second)
}
