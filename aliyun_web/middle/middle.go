package middle

import (
	"github.com/gin-gonic/gin"
	"main/aliyun_web/global"
	"sync/atomic"
)

func IP_exist(ctx *gin.Context) {
	if len(global.Server_ip) == 0 {
		ctx.JSON(502, gin.H{"msg": "上游服务器出错"})
		ctx.Abort()
		return
	}
	if atomic.LoadInt64(&global.Server_ip_Index) >= int64(len(global.Server_ip)) {
		atomic.StoreInt64(&global.Server_ip_Index, 0)
	}
	ctx.Set("Server_ip_Index", atomic.LoadInt64(&global.Server_ip_Index))
	atomic.AddInt64(&global.Server_ip_Index, 1)
	ctx.Next()
}
