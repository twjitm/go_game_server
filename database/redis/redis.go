package redis

import (
	"github.com/go-redis/redis"
	"time"
)

var redisClient *Client

type Client struct {
	Host  string
	Port  int
	Redis *redis.Client
}

func init() {
	NewClient()
}

func NewClient() *Client {
	if redisClient == nil {
		client := redis.NewClient(&redis.Options{
			Addr:         "127.0.0.1:6379",
			DialTimeout:  time.Second,
			ReadTimeout:  time.Second,
			WriteTimeout: time.Second,
		})
		redisClient = &Client{
			Host:  "127.0.0.1",
			Port:  6379,
			Redis: client,
		}
	}
	return redisClient
}

func get(key string) interface{} {
	stringCmd := NewClient().Redis.Get(key)
	val := stringCmd.Val()
	return val
}

func getBitInt(key string) int64 {

}
