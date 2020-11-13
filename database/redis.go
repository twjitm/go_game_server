package database

import (
"github.com/go-redis/redis"
"go_game_server/conf"
"time"
)

var redisClient map[string]*redis.Client

func Redis(shard string) *redis.Client {
	client := redisClient[shard]
	config := conf.GetRedisPool(shard)
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:         config.Host,
			DialTimeout:  time.Second,
			ReadTimeout:  time.Second,
			WriteTimeout: time.Second,
		})
		redisClient[shard] = client
	}
	return client
}
