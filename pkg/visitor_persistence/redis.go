package visitor_persistence

import "github.com/gomodule/redigo/redis"

type Redis struct {
	Con *redis.Conn
}

func (r *Redis) Add(id string) error {
	return nil
}

func (r *Redis) List(ids []string) ([]map[int64]int64, error) {
	return nil, nil
}
