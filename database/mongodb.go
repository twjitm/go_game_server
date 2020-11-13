package database

import (
	"context"
	"encoding/json"
	"github.com/prometheus/common/log"
	"go_game_server/conf"
	"go_game_server/cutil"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

type MgoNames struct {
	database   string
	collection string
}

func GetNamespace(collection string) *MgoNames {
	mn := &MgoNames{
		collection: collection,
	}
	if len(conf.ServerConf.MongoDBConf.Suffix) > 0 {
		mn.database = "gs_server_" + conf.ServerConf.MongoDBConf.Suffix
	} else {
		mn.database = "gs_server_user"
	}
	return mn
}

func SetupMongoDB() error {
	if !conf.ServerConf.MongoDBConf.InUse {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := conf.ServerConf.MongoDBConf.URI
	if conf.ServerConf.MongoDBConf.Sharding {
		uri = conf.ServerConf.MongoDBConf.MongosURI
	}
	maxPoolSize := uint64(conf.ServerConf.MongoDBConf.MaxPoolSize)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(maxPoolSize))
	if err != nil {
		log.Error(err)
		return err
	}
	MongoClient = client
	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Error(err)
	}
	return err
}

func UpdateOne(mgo *MgoNames, filter interface{}, data interface{}) error {
	if !conf.ServerConf.MongoDBConf.InUse {
		return nil
	}
	collection := MongoClient.Database(mgo.database).Collection(mgo.collection)
	if collection == nil {
		return nil
	}
	opts := options.Update()
	opts.SetUpsert(true)
	value := bson.M{"$set": data}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.ServerConf.MongoDBConf.TimeOut)*time.Second)
	defer cancel()
	_, err := collection.UpdateOne(ctx, filter, value, opts)
	if err != nil {
		log.Error(err)
	}
	return err
}

func FindAll(mgo *MgoNames, filter interface{}, toData func(ctx context.Context, cur *mongo.Cursor) bool) error {
	if !conf.ServerConf.MongoDBConf.InUse {
		return nil
	}
	collection := MongoClient.Database(mgo.database).Collection(mgo.collection)
	if collection == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second) // reload所有kingdom比较耗时，这里使用一个比较大的超时时间
	defer cancel()
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error(err)
		return err
	}
	if !toData(ctx, cur) {
		return nil
	}
	err = cur.Close(ctx)
	return err
}

func FindOne(mgo *MgoNames, filter interface{}, toData func(ctx context.Context, result *mongo.SingleResult) (successful bool, err error)) error {
	if !conf.ServerConf.MongoDBConf.InUse {
		return nil
	}
	collection := MongoClient.Database(mgo.database).Collection(mgo.collection)
	if collection == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.ServerConf.MongoDBConf.TimeOut)*time.Second)
	defer cancel()
	singleResult := collection.FindOne(ctx, filter)
	if singleResult != nil {
		_, err := toData(ctx, singleResult)
		return err
	}
	return nil
}

func DeleteMany(mgo *MgoNames, filter interface{}) error {
	if !conf.ServerConf.MongoDBConf.InUse {
		return nil
	}
	collection := MongoClient.Database(mgo.database).Collection(mgo.collection)
	if collection == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.ServerConf.MongoDBConf.TimeOut)*time.Second)
	defer cancel()
	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Error(err)
	}
	return err
}

func InsertOne(mgo *MgoNames, value interface{}) error {
	if !conf.ServerConf.MongoDBConf.InUse {
		return nil
	}
	collection := MongoClient.Database(mgo.database).Collection(mgo.collection)
	if collection == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.ServerConf.MongoDBConf.TimeOut)*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, value)
	return err
}

func InsertMany(mgo *MgoNames, values []interface{}) error {
	if !conf.ServerConf.MongoDBConf.InUse {
		return nil
	}
	collection := MongoClient.Database(mgo.database).Collection(mgo.collection)
	if collection == nil {
		return nil
	}
	opts := options.InsertMany()
	opts.SetOrdered(false)
	perRows := int32(50000)
	var err error
	s1 := time.Now()
	if int32(len(values)) > perRows {
		perNum := int32(len(values)) / perRows
		for i := int32(0); i < perNum+1; i++ {
			startIdx := i * perRows
			endIdx := cutil.Min32(startIdx+perRows, int32(len(values)))
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.ServerConf.MongoDBConf.TimeOut)*time.Second)
			_, err = collection.InsertMany(ctx, values[startIdx:endIdx], opts)
			cancel()
			if err != nil {
				log.Error(err)
				return err
			}
		}
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.ServerConf.MongoDBConf.TimeOut)*time.Second)
		defer cancel()
		_, err = collection.InsertMany(ctx, values, opts)
		if err != nil {
			log.Error(err)
		}
	}
	et := time.Now().Sub(s1)
	log.Info("insert docs counter: ", len(values))
	log.Info("insert elapsed time: ", et)
	return err
}

func BatchExecute(mgo *MgoNames, bulkModels []mongo.WriteModel) error {
	if !conf.ServerConf.MongoDBConf.InUse {
		return nil
	}
	if len(bulkModels) == 0 {
		return nil
	} // 没有需要持久化的内容
	collection := MongoClient.Database(mgo.database).Collection(mgo.collection)
	if collection == nil {
		return nil
	}
	opt := options.BulkWrite()
	opt.SetOrdered(true)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.ServerConf.MongoDBConf.TimeOut)*time.Second)
	defer cancel()
	_, err := collection.BulkWrite(ctx, bulkModels, opt)
	if err != nil {
		log.Error(err)
	}
	return err
}

func CopyUpdateValues(v1, v2 interface{}) bool {
	vv1, vv2 := reflect.ValueOf(v1), reflect.ValueOf(v2)
	switch vv2.Kind() {
	case reflect.Array:
		for i := 0; i < vv2.Len(); i++ {
			vv1 = reflect.Append(vv1, vv2.Index(i))
		}
	case reflect.Slice:
		for i := 0; i < vv2.Len(); i++ {
			vv1 = reflect.Append(vv1, vv2.Index(i))
		}
	case reflect.Ptr:
		return CopyUpdateValues(vv1.Elem(), vv2.Elem())
	case reflect.Struct:
		//for i, n := 0, vv1.NumField(); i < n; i++ {
		//	CopyUpdateValues(vv1.Field(i).Elem(), vv2.Field(i).Elem())
		//}
		vv1 = vv2
	case reflect.Map:
		for _, k := range vv1.MapKeys() {
			val1 := vv1.MapIndex(k)
			val2 := vv2.MapIndex(k)
			CopyUpdateValues(val1, val2)
		}
	default:
		vv1.Elem().Set(vv2.Elem())
	}
	return true
}

func DiffUpdateFields(u1, u2 interface{}) interface{} {
	if u1 == nil {
		return u2
	}
	d, _ := json.Marshal(u2)
	var fields map[string]interface{}
	json.Unmarshal(d, &fields)
	v1 := reflect.ValueOf(u1)
	v2 := reflect.ValueOf(u2)
	t2 := reflect.TypeOf(u2)
	if v1.NumField() != v2.NumField() {
		return fields
	}
	for i := 0; i < v1.NumField(); i++ {
		fieldName := t2.Field(i).Tag.Get("json")
		vv1 := v1.Field(i).Interface()
		vv2 := v2.Field(i).Interface()
		if reflect.DeepEqual(vv1, vv2) {
			delete(fields, fieldName)
		}
	}
	if len(fields) == 0 { // 没有差异时返回nil，便于后续判断
		return nil
	}
	return fields
}
