package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"sync"
	"time"
)

//etcd 学习

type clientMgr struct {
	client     *clientv3.Client
	serverList map[string]string
	lock       *sync.Mutex
}

var EtcdClient clientMgr

func init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second,
	})
	if err != nil {
		defer func() {
			if cli != nil {
				_ = cli.Close()
			}
		}()
	}
	if cli == nil {
		fmt.Println("etcd server not connection")
		return

	}
	EtcdClient = clientMgr{client: cli}

}

func (client *clientMgr) Put(key string, value interface{}) bool {

	var data []byte
	_ = json.Unmarshal(data, value)
	result, err := client.client.Put(context.Background(), key, string(data))

	if err != nil {
		fmt.Printf("send key=%s filed ,please check\n", key)
	}
	if result == nil {

		return false
	}
	return true

}
func (client *clientMgr) Get(key string, toData func(data string)) interface{} {
	result, err := client.client.Get(context.Background(), key)
	if err != nil {
	}

	if result == nil {
		return ""
	}
	kvs := result.Kvs
	if len(kvs) == 0 {
		fmt.Printf("not found key=%s\n", key)
	}
	for i := 0; i < len(kvs); i++ {
		v := kvs[i].Value
		toData(string(v))
	}
	return nil
}
func (client *clientMgr) PutMember(values []interface{}) {

	var fileds []string
	for i := 0; i < len(values); i++ {
		data, err := json.Marshal(values[i])
		if err != nil {
			continue
		}
		fileds[i] = string(data)
	}
	_, _ = client.client.MemberAdd(context.Background(), fileds)
}

func (client *clientMgr) Watcher(key string, do func(event *clientv3.Event)) {
	go func() {
		watch := EtcdClient.client.Watch(context.Background(), key)
		for result := range watch {
			EtcdClient.lock.Lock()
			fmt.Println("watch key is change")
			events := result.Events
			for i := 0; i < len(events); i++ {
				event := events[i]
				fmt.Println(event.Type.String())
				//事件类型
				do(event)
				switch event.Type {
				case mvccpb.PUT:
					EtcdClient.serverList[string(event.Kv.Key)] = string(event.Kv.Value)
				case mvccpb.DELETE:
					delete(EtcdClient.serverList, string(event.Kv.Key))
				}
			}
			EtcdClient.lock.Lock()
		}

		//select {
		//case watchResponse := <-watch:
		//	fmt.Println("watch key is change")
		//	events := watchResponse.Events
		//	for i := 0; i < len(events); i++ {
		//		event := events[i]
		//		//事件类型
		//		switch event.Type {
		//		case mvccpb.PUT:
		//			//新增节点 todo
		//		case mvccpb.DELETE:
		//			//删除节点 todo
		//		}
		//	}
		//}
	}()
}
