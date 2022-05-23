package global

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
)

var Control_ip string

var Host string

var Root *gin.Engine

var Server_ip []string

var Server_ip_Index int64

func init() {
	Root = gin.New()

	h, _ := os.Getwd()
	viper.AddConfigPath(h + `\aliyun\static`)
	viper.SetConfigName("setting")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	t := viper.Get("control_ip")
	if t == nil {
		panic(errors.New("配置文件中无control_ip"))
	}
	Control_ip = t.(string)
}

func Init() {
	li, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	Host = strconv.Itoa(li.Addr().(*net.TCPAddr).Port)
	li.Close()

	re, err := http.Get("http://" + Control_ip + "/get_server?host=" + Host)
	if err != nil {
		panic(err)
	}
	m := map[string]interface{}{}
	b, _ := ioutil.ReadAll(re.Body)
	re.Body.Close()
	json.Unmarshal(b, &m)
	if m["Server_ip"] == nil {
		panic(errors.New("接收服务器ip出错"))
	}
	Server_ip = m["Server_ip"].([]string)
}
