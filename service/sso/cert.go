package sso

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go_game_server/database"
)

/**应用认证发布 任何*/
//任何应用必须认证
type Cert struct {
	AppId  string `json:"app_id"`
	Token  string `json:"token"`
	Status int8   `json:"status"`
}

type CertServiceImpl struct {
}

var CertService CertServiceImpl

func (c CertServiceImpl) GetAll() *[]Cert {
	var certs *[]Cert
	database.FindAll(database.GetNamespace("cert"), bson.M{
		"status": 1,
	}, func(ctx context.Context, cur *mongo.Cursor) bool {
		err := cur.Decode(certs)
		return err == nil
	})
	return certs
}

func (c CertServiceImpl) GetById(appId string) *Cert {
	var cert *Cert
	database.FindOne(database.GetNamespace("cert"), bson.M{"app_id": appId},
		func(ctx context.Context, result *mongo.SingleResult) (bool, error) {
			err := result.Decode(cert)
			if err != nil {
				fmt.Println("get cert error")
				return false, err
			}
			return true, nil
		})
	return cert
}

func (c CertServiceImpl) Add(cert *Cert) error {
	if cert.AppId == "" || cert.Token == "" {
		return fmt.Errorf("data error")
	}
	if c.GetById(cert.AppId) != nil {
		return fmt.Errorf("cert exists")
	}
	err := database.InsertOne(database.GetNamespace("cert"), cert)
	return err
}

func (c CertServiceImpl) Delete(appId string) error {
	return database.DeleteMany(database.GetNamespace("cert"), bson.M{
		"app_id": appId,
	})
}

func (c CertServiceImpl) Valid(appId, token string) bool {
	cert := c.GetById(appId)
	if cert == nil {
		return false
	}
	return cert.Token == token
}

