package minesweepersvc

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"os"
)

type DB interface {
	Save(key string, value interface{}) error
	Get(key string) (*User, error)
	GetGame(key string) (*Game, error)
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
	uJson, _ := json.Marshal(value)
	_, err := r.Do("SET", key, uJson)
	return err
}

func (r *RedisDB) Get(key string) (*User, error) {
	reply, err := redis.String(r.Do("GET", key))
	logs.Sugar().Infof("%s", reply)
	var user User
	err = json.Unmarshal([]byte(reply), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RedisDB) GetGame(key string) (*Game, error) {
	reply, err := redis.String(r.Do("GET", key))
	logs.Sugar().Infof("%s", reply)
	var game Game
	err = json.Unmarshal([]byte(reply), &game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func getConn() redis.Conn {
	redisURL := os.Getenv("REDISCLOUD_URL")
	logs.Log().Info(redisURL)
	c, err := redis.DialURL(redisURL)
	if err != nil {
		logs.Log().Fatal("cannot connect with Redis")
		panic(err)
	}
	return c
}
