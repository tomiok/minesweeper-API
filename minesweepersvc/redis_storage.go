package minesweepersvc

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"go.uber.org/zap"
	"os"
)

type RedisDB struct {
	redis.Conn
}

func NewDB() DB {
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

	if err != nil {
		return nil, err
	}
	logs.Sugar().Info(reply)
	var user User
	err = json.Unmarshal([]byte(reply), &user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RedisDB) GetGame(key string) (*Game, error) {
	reply, err := redis.String(r.Do("GET", key))

	if err != nil {
		return nil, err
	}
	logs.Sugar().Info(reply)
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

func (r *RedisDB) FlushAll() error {
	s, err := redis.String(r.Do("FLUSHALL"))
	logs.Log().Info("flushed!", zap.String("reply", s))
	return err
}

func getConn() redis.Conn {
	logs.Log().Info("connecting redis...")
	redisURL := os.Getenv("REDISCLOUD_URL")

	c, err := connectRedis(redisURL)

	if err != nil {
		logs.Log().Fatal("cannot connect with Redis")
		panic(err)
	}

	return c
}

//format url ==> "redis://redis:6379"
func connectRedis(redisURL string) (redis.Conn, error) {
	if redisURL != "" {
		redisPassword := os.Getenv("REDISCLOUD_PASSWORD")
		return redis.DialURL(redisURL, redis.DialPassword(redisPassword))
	} else {
		localURL := os.Getenv("REDIS_LOCAL_URL")
		redisURL = fmt.Sprintf("redis://%s", localURL)
		return redis.DialURL(redisURL)
	}
}
