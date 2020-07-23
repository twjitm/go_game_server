package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type RedisPoll struct {
	Host    string
	TimeOut int64
}

func init() {

}

func GetRedisConfig() *[]RedisPoll {
	config := viper.New()

	config.SetConfigFile("redis")
	config.SetConfigType("json")
	config.AddConfigPath("conf/")
	config.AddConfigPath("$GOPATH/src/")

	var pools = []RedisPoll{}
	keys := config.AllKeys()

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Print(keys)
	err := config.Unmarshal(&pools)
	if err != nil {
		fmt.Errorf("redis config error")
	}
	return &pools
}
