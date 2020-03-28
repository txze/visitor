package redis_client

import "github.com/gomodule/redigo/redis"

var MasterPool *redis.Pool
var SalvePool *redis.Pool

func GetMasterCon() redis.Conn {
	return MasterPool.Get()
}

func GetSalveCon() redis.Conn {
	return SalvePool.Get()
}
