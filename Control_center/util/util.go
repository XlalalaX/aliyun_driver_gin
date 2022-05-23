package util

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"main/Control_center/global"
	"os"
)

func Save_ip() {
	f, err := os.Create("ip+host.txt")
	b, err := json.Marshal(&global.M_ip)
	if err != nil {
		panic(errors.New("保存ip失败：" + err.Error()))
	}
	f.Write(b)
	f.Close()
}

func Delete_ex_timeout_file() {
	msgs := global.Rdb.PSubscribe("__keyevent@1__:expired").Channel()

	for {
		msg := <-msgs
		re, err := global.Aliyun.SearchNameInFolder(global.User, msg.Payload, global.Root_path)
		if err != nil || re.Code != "" || len(re.Items) == 0 {
			log.Println(err)
			continue
		}
		resp, err := global.Aliyun.RemoveFile(global.User, re.Items[0].FileId)
		if err != nil || resp.Code != "" {
			log.Println(err)
			log.Println(resp.Message)
			continue
		}
	}
}
