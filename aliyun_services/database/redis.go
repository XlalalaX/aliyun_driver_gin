package database

import "main/aliyun_services/global"

func Delet_redis_string(string string) (err error) {

	_, err = global.Rdb.Del(string).Result()
	return

}
