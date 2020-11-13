package sso

import (
	"go_game_server/database"
	"go_game_server/service/base"
)

/**应用认证发布 任何*/
//任何应用必须认证
type Cert struct {
	AppId  string `json:"app_id"`
	Token  string `json:"token"`
	Status int8   `json:"status"`
}

type CertServiceImpl struct {
	base.MongodbService
}

var GertService CertServiceImpl

func (c CertServiceImpl) GetAll() []interface{} {

	database.GetApisListByIds()
	return nil
}

func (c CertServiceImpl) GetById(id interface{}) interface{} {
	return nil
}

