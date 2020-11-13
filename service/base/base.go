package base

type MongodbService interface {
	GetById(id interface{}) interface{}
	GetAll() []interface{}
}
