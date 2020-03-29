package visitor_persistence

import (
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"visitor/client/redis_client"
)

func TestRedis_Add(t *testing.T) {
	con := redis_client.GetMasterCon()
	defer con.Close()
	id := "testid"
	con.Do("DEL", RedisVisitorKeyPre+id)
	r := Redis{&con}
	err := r.Add(id, 1)
	if err != nil {
		t.Error(err)
		return
	}
	err = r.Add(id, 2)
	if err != nil {
		t.Error(err)
		return
	}
	err = r.Add(id, int64(time.Minute)+1)
	if err != nil {
		t.Error(err)
		return
	}
	err = r.Add(id, int64(time.Minute*2)+1)
	if err != nil {
		t.Error(err)
		return
	}
	reply, err := r.List(id, []int64{1, 2, 3}, int64(time.Minute)*3)
	if err != nil {
		t.Error(err)
		return
	}
	if reply[1] != 1 {
		t.Error(1)
		return
	}
	if reply[2] != 2 {
		t.Error(2)
		return
	}
	if reply[3] != 4 {
		t.Error(3)
		return
	}
}

func init() {
	//初始化redis
	redis_client.MasterPool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	redis_client.SalvePool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
