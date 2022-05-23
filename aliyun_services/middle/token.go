package middle

import (
	"github.com/gin-gonic/gin"
	"main/aliyun_services/global"
)

func Token_exist_mid(ctx *gin.Context) {
	if !global.Token_exist {
		ctx.JSON(500, gin.H{"msg": "服务器出错"})
		ctx.Abort()
		return
	}
	ctx.Next()
}
