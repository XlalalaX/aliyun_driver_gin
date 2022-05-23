package global

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jakeslee/aliyundrive"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
)

var Root *gin.Engine

//serveer端ip
var M_ip map[string]int

//web端ip
var Web_ip map[string]int

var User *aliyundrive.Credential

var Aliyun *aliyundrive.AliyunDrive

var Rdb *redis.Client

var Root_path string

var Port string

var Db_ip string

func init() {
	Root = gin.New()
	M_ip = map[string]int{}
	h, _ := os.Getwd()
	fmt.Printf(h)
	viper.SetConfigName("token")                      // name of config file (without extension)
	viper.SetConfigType("yaml")                       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(h + `\Control_center\static`) // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("读取配置文件失败: %w \n", err))
	}
}

func Init() {
	token_t := viper.Get("token")
	root_path_t := viper.Get("root_path")
	db_ip_t := viper.Get("db_ip")
	if token_t == nil || root_path_t == nil || db_ip_t == nil {
		panic(errors.New("读取本地配置文件失败"))
	}
	token := token_t.(string)
	Root_path = root_path_t.(string)
	Db_ip = db_ip_t.(string)

	if token == "" || Root_path == "" {
		panic(errors.New("读取本地配置文件失败"))
	}

	User = &aliyundrive.Credential{RefreshToken: token}
	Aliyun = aliyundrive.NewClient(&aliyundrive.Options{AutoRefresh: true})
	var err error
	User, err = Aliyun.AddCredential(aliyundrive.NewCredential(User))
	if err != nil {
		panic(err)
	}
	Aliyun.RefreshAllToken()
	log.Println("用户名:", User.Name)

	re, err := Aliyun.SearchNameInFolder(User, Root_path, "root")
	if err != nil {
		panic(err)
	} else {
		Root_path = re.Items[0].FileId
	}

	log.Println("root_id:", Root_path)

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})
	err = Rdb.Ping().Err()
	if err != nil {
		panic(err)
	}

	recover_ip()
}

func recover_ip() {
	f, err := os.Open("ip+host.txt")
	if err != nil {
		log.Println("无历史ip文件")
		return
	}

	b, _ := ioutil.ReadAll(f)
	json.Unmarshal(b, &M_ip)
}
