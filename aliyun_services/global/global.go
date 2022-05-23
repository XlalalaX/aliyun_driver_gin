package global

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jakeslee/aliyundrive"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var Token_exist bool

var Root *gin.Engine

var Aliyun *aliyundrive.AliyunDrive

var User *aliyundrive.Credential

var Sava_path_Id string

var Rdb *redis.Client

var Rdb_2 *redis.Client

var m map[string]interface{}

var token string

var Port string

var db_ip string

var Debug bool

var Control_ip string

func Init() {
	var err error
	Token_exist = true

	rand.Seed(time.Now().Unix())
	h, _ := os.Getwd()
	viper.AddConfigPath(h + `\aliyun_services\static`)
	viper.SetConfigName("ip_host")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(errors.New("读取本地配置失败:" + err.Error()))
	}
	d := viper.Get("debug")
	if d != nil {
		Debug = d.(bool)
		Port = viper.Get("port").(string)
	}

	control_ip := viper.Get("control_ip")
	if control_ip == nil {
		panic(errors.New("配置文件中无控制中心ip"))
	}
	Control_ip = control_ip.(string)

	if Port == "" {
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			panic(err)
		}

		fmt.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)
		Port = strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
		listener.Close()
	}

	Root = gin.New()
	get_config()

	User = &aliyundrive.Credential{RefreshToken: token}
	Aliyun = aliyundrive.NewClient(&aliyundrive.Options{AutoRefresh: true})
	User, _ = Aliyun.AddCredential(aliyundrive.NewCredential(User))
	go func() {
		for {
			_, err := Aliyun.RefreshToken(User)
			log.Println("user:", User.Name)
			if err != nil {
				for {
					err = get_config()
					if err != nil {
						Token_exist = false
						continue
					} else {
						Token_exist = true
					}
					User = &aliyundrive.Credential{RefreshToken: token}
					Aliyun = aliyundrive.NewClient(&aliyundrive.Options{AutoRefresh: true})

					User, err = Aliyun.AddCredential(aliyundrive.NewCredential(User))
					if err != nil {
						Token_exist = false
					} else {
						Token_exist = true
						break
					}
					time.Sleep(time.Second * 10)
				}
			}
			time.Sleep(time.Hour)
		}
	}()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})

	err = Rdb.Ping().Err()
	if err != nil {
		panic(err)
	}

	Rdb_2 = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       1,  // 数据库
		PoolSize: 20, // 连接池大小
	})

	err = Rdb_2.Ping().Err()
	if err != nil {
		panic(err)
	}
}

func get_config() (err error) {

	resp, err := http.Get("http://" + Control_ip + "/add_server" + "?host=" + Port)
	if err != nil {
		return errors.New("从控制中心拉取token失败" + err.Error())
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	m = map[string]interface{}{}
	json.Unmarshal(b, &m)

	if m["token"] == nil || m["token"].(string) == "" {
		return errors.New("拉取token失败")
	}
	if m["root_path"] == nil || m["root_path"].(string) == "" {
		return errors.New("拉取root_path失败")
	}
	if m["db_ip"] == nil || m["db_ip"].(string) == "" {
		return errors.New("拉取root_path失败")
	}
	token = m["token"].(string)
	Sava_path_Id = m["root_path"].(string)
	db_ip = m["db_ip"].(string)
	return nil
}

func Reflash_token(ctx *gin.Context) {
	token = ctx.PostForm("token")
	Sava_path_Id = ctx.PostForm("root_path")

	User = &aliyundrive.Credential{RefreshToken: token}
	Aliyun = aliyundrive.NewClient(&aliyundrive.Options{AutoRefresh: true})
	var err error
	User, err = Aliyun.AddCredential(aliyundrive.NewCredential(User))
	if err != nil {
		Token_exist = false
	} else {
		Token_exist = true
	}
	ctx.JSON(200, gin.H{"msg": "刷新token成功"})
}
