package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

type RedisPoll struct {
	Host    string `json:"host"`
	TimeOut int64  `json:"timeout"`
	Shard   string `json:"shard"`
}
type Config struct {
	Pools []RedisPoll `json:"pools"`
}

var redisPools []RedisPoll

func init() {
	fmt.Println("configMgr init")
	GetRedisConfig()
}

func GetRedisPool(shard string) (pool *RedisPoll) {
	for ps := range redisPools {
		pp := redisPools[ps]
		if pp.Shard == shard {
			pool = &pp
		}
	}
	return pool
}

func GetRedisConfig() *[]RedisPoll {

	globalViper := viper.New()
	configPath := filepath.Join("conf", "redis.json")
	globalViper.SetConfigFile(configPath)
	var pools Config
	if err := globalViper.ReadInConfig(); err != nil {
		panic(err)
	}
	err := globalViper.Unmarshal(&pools)
	if err != nil {
		panic(err)
	}
	fmt.Println(pools)
	redisPools = pools.Pools
	return &pools.Pools
}
