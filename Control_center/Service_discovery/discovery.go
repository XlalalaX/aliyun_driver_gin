package Service_discovery

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/Control_center/global"
	"net/http"
	"strings"
)

//server服务发现，注册
func Discovery(ctx *gin.Context) {
	ip := ctx.ClientIP()
	host := ctx.Query("host")
	ip = ip + ":" + host
	if host == "" {
		ctx.JSON(400, gin.H{"msg": "缺少host"})
		ctx.Abort()
		return
	}
	global.M_ip[ip] = 1
	log.Println("add_server_ip:", ip)
	go Notify_web_server_ip()
	ctx.JSON(200, gin.H{"token": global.User.RefreshToken, "root_path": global.Root_path, "db_ip": global.Db_ip})

}

//通知web端改变后的server_ip
func Notify_web_server_ip() {
	ss := []string{}

	for i, _ := range global.M_ip {
		ss = append(ss, i)
	}

	c := http.Client{}

	for i, _ := range global.Web_ip {
		r, err := http.NewRequest("POST", "http://"+i, nil)
		if err != nil {
			global.Web_ip[i]++
		}
		r.PostForm.Set("Server_ip", strings.Join(ss, ","))
		re, err := c.Do(r)
		if err != nil || re.StatusCode != http.StatusOK {
			global.Web_ip[i]++
		}
		if global.Web_ip[i] >= 6 {
			delete(global.Web_ip, i)
		}
	}

}

//web端
func Get_server_ip(ctx *gin.Context) {
	ip := ctx.ClientIP()
	host := ctx.Query("host")
	ip = ip + ":" + host
	if host == "" {
		ctx.JSON(400, gin.H{"msg": "缺少host"})
		ctx.Abort()
		return
	}
	global.Web_ip[ip] = 1
	log.Println("add_web_ip:", ip)
	ss := []string{}

	for i, _ := range global.M_ip {
		ss = append(ss, i)
	}
	ctx.JSON(200, gin.H{"Server_ip": ss})
}

//更新server端token
func Reflash_token() {
	c := http.Client{}

	for s, _ := range global.M_ip {
		r, err := http.NewRequest("POST", "http://"+s+"/reflash_token", nil)
		if err != nil {
			global.M_ip[s]++
		}
		r.Header.Set("token", global.User.RefreshToken)
		r.Header.Set("root_path", global.Root_path)
		re, err := c.Do(r)
		if err != nil || re.StatusCode != http.StatusOK {
			global.M_ip[s]++
		}
	}
}
