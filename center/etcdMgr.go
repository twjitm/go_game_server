package center

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	_ "log"
	"time"
)

//etcd 学习

type clientMgr struct {
	client *clientv3.Client
}

var globalClient clientMgr

func init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0≥0.1:2379"},
		DialTimeout: time.Second,
	})
	if err != nil {

		defer cli.Close()
	}
	globalClient = clientMgr{client: cli}

}

func (client *clientMgr) Put(key string, value interface{}) bool {

	var data []byte
	_ = json.Unmarshal(data, value)
	result, err := client.client.Put(context.Background(), key, string(data))

	if err != nil {
		fmt.Println("send key=%s filed ,please check", key)
	}
	if result == nil {

		return false
	}
	return true

}
func (client *clientMgr) Get(key string) interface{} {
	result, err := client.client.Get(context.Background(), key)
	if err != nil {
	}

	kvs := result.Kvs
	for i := 0; i < len(kvs); i++ {


	}

}
