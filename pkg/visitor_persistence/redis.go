package visitor_persistence

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
	"visitor/client/redis_client"
)

const (
	ZADD   = "ZADD"
	ZCOUNT = "ZCOUNT"
)
const RedisVisitorKeyPre = "visitor"

type Redis struct {
	Con *redis.Conn
}

//目前先支持单个用户的添加和列出
func (r *Redis) Add(id string, nano int64) error {
	var con = redis_client.GetMasterCon()
	defer con.Close()
	item := fmt.Sprintf("%s%d", id, nano) //使用这个作为唯一有一个问题 极端情况下 用户在同一纳秒登录好多次只能统计到一次 这种情况可以根据下面被忽略的参数捕获
	_, err := redis.Int(con.Do(ZADD, RedisVisitorKeyPre+id, nano, item))
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) List(id string, span []int64, nano int64) (map[int64]int64, error) {
	var con = redis_client.GetMasterCon()
	defer con.Close()
	//span里面的单位是分钟
	var result = map[int64]int64{}
	for _, item := range span {
		min := nano - item*int64(time.Minute)
		count, err := redis.Int64(con.Do(ZCOUNT, RedisVisitorKeyPre+id, min, nano))
		if err != nil {
			return nil, err
		}
		result[item] = count
	}
	return result, nil
}
