package minesweepersvc

import (
	"github.com/gomodule/redigo/redis"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"os"
)

type DB interface {
	Save(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

type RedisDB struct {
	redis.Conn
}

func New() DB {
	return &RedisDB{
		getConn(),
	}
}

func (r *RedisDB) Save(key string, value interface{}) error {
	return r.Send("SET", key, value)
}

func (r *RedisDB) Get(key string) (interface{}, error) {
	_ = r.Send("GET", key)

	return r.Receive() // reply from GET
}

func getConn() redis.Conn {
	redisURL := os.Getenv("REDISCLOUD_URL")
	logs.Log().Info(redisURL)
	c, err := redis.DialURL(redisURL)
	if err != nil {
		logs.Log().Fatal("cannot connect with Redis")
		panic(err)
	}
	err = c.Send("PING")

	if err != nil {
		logs.Log().Fatal("cannot connect with Redis")
		panic(err)
	}
	return c
}
//redis://rediscloud:DTtY29OIKVIk3zDsWsTuSoyZhdFErc6W@redis-12571.c8.us-east-1-4.ec2.cloud.redislabs.com:12571C