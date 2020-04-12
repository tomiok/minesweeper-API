package minesweepersvc

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"go.uber.org/zap"
	"os"
)

type DB interface {
	Save(key string, value interface{}) error
	Get(key string) (*User, error)
	GetGame(key string) (*Game, error)
	Exists(key string) bool
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
	uJson, err := json.Marshal(value)

	if err != nil {
		logs.Log().Error("cannot marshal current structure", zap.Error(err))
		return err
	}
	_, err = r.Do("SET", key, uJson)
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

func (r *RedisDB) Exists(key string) bool {
	reply, err := redis.Int(r.Do("EXISTS", key))

	if err != nil {
		logs.Log().Error("cannot fetch value with exists command")
		return false
	}

	return reply > 0
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
