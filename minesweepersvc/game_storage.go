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
	c, err := redis.DialURL(os.Getenv("REDIS_URL"))
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
