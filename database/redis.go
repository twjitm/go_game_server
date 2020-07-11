package database

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

func GetRedisClient() (redisClient *redis.Client) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	return client
}

type UserInfo struct {
	ID       int64 `json:"ID"`
	Name     string `json:"Name"`
	Birthday string `json:"Birthday"`
}

func (user *UserInfo) toString() string {
	buf, err := json.Marshal(user)
	if err != nil {
		return ""
	}
	return string(buf)
}

func SaveUser(info UserInfo) bool {
	client := GetRedisClient()
	hKey := string(info.ID)
	bol := client.HSet(UNFRIENDLIEST, hKey, info)
	if !bol.Val() {
		panic("save exception")
	}
	return bol.Val()
}

func GetUserById(id int64) UserInfo {
	client := GetRedisClient()
	hKey := string(id)
	result := client.HGet(UNFRIENDLIEST, hKey)
	info := UserInfo{}
	_ = json.Unmarshal([]byte(result.Val()), &info)

}

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(UNFRIENDLIEST, "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
