package conf

import (
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
)

type MongodbConf struct {
	Suffix      string
	MaxPoolSize int32
	URI         string
	InUse       bool
	Sharding    bool
	MongosURI   string
	TimeOut     int64
}


type GlobalServerConf struct {
	MongoDBConf *MongodbConf
}

var ServerConf GlobalServerConf

func init() {
	ServerConf = GlobalServerConf{
		MongoDBConf: nil,
	}
	data, err := ioutil.ReadFile("conf/mongodb.json")
	mongo := MongodbConf{}
	err = json2.Unmarshal(data, &mongo)
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	ServerConf.MongoDBConf = &mongo

}
