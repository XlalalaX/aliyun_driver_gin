package action

import (
	"github.com/gin-gonic/gin"
	"main/aliyun_web/global"
	"strings"
)

func Reflash_server_ip(ctx *gin.Context) {
	s := ctx.PostForm("Server_ip")
	global.Server_ip = strings.Split(s, ",")
	global.Server_ip_Index = 0
	ctx.JSON(200, gin.H{"mag": "更新serverip成功"})
}

func Get_ip(ctx *gin.Context) {
	index, _ := ctx.Get("Server_ip_Index")
	s := global.Server_ip[index.(int)]

	ctx.JSON(200, gin.H{"Server_ip": s})
}
