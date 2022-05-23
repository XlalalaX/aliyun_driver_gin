package aliyun_web

import (
	"main/aliyun_web/action"
	"main/aliyun_web/global"
	"main/aliyun_web/middle"
)

func main() {
	global.Init()

	global.Root.POST("/reflash_server_ip", action.Reflash_server_ip)

	users := global.Root.Group("/user")
	users.Use(middle.IP_exist)

	users.GET("/get_ip", action.Get_ip)

	global.Root.Run(":" + global.Host)

}
