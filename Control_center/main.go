package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"main/Control_center/Service_discovery"
	"main/Control_center/global"
	"main/Control_center/heartbeat"
	"main/Control_center/util"
	"time"
)

func main() {
	global.Init()
	viper.OnConfigChange(func(in fsnotify.Event) {
		global.Init()
		Service_discovery.Reflash_token()
	})

	viper.WatchConfig()

	go func() {
		for {
			util.Save_ip()
			time.Sleep(time.Second * 60)
		}
	}()

	go func() {
		for {
			heartbeat.Hearbeat()
			time.Sleep(time.Second * 5)
		}
	}()

	go func() {
		util.Delete_ex_timeout_file()
	}()

	global.Root.GET("/add_server", Service_discovery.Discovery)
	global.Root.GET("/get_server", Service_discovery.Get_server_ip)

	global.Root.Run(":9998")
}
