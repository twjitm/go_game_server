package sso

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go_game_server/common"
	"go_game_server/database"
	"time"
)

const SESSION_TIME = 1800

type SSoPo struct {
	SessionId string `json:"session_id"`
	Uid       int64  `json:"uid"`
	Token     string `json:"token"`
	Mtime     int64  `json:"mtime"`
}

func Register(po SSoPo) *common.ResultPo {
	result := common.CreateResultPo()
	redis := database.Redis("sso")
	//old session
	oldSessionId := redis.HGet(database.SESSION_POOL, string(po.Uid)).String()
	if oldSessionId != "" {
		redis.HGetAll(database.GetSSoSessionKey(oldSessionId))
	}
	return result
}

func Validate(sessionId, token string) *common.ResultPo {
	result := common.CreateResultPo()
	redis := database.Redis("sso")
	key := database.GetSSoSessionKey(sessionId)

	command := redis.HGetAll(key)
	data := command.Val()
	if data == nil {
		result.Code = common.ERROR_SESSION_NOTFOUND
		result.Message = "session validate error : not found"
		return result
	}
	sso := SSoPo{}
	err := mapstructure.Decode(data, &sso)
	if err != nil {
		result.Code = common.ERROR_SESSION_NOTFOUND
		result.Message = "session  decode err"
		return result
	}
	now := time.Now()
	//session expire
	if sso.Mtime+SESSION_TIME > now.Unix() {
		result.Code = common.ERROR_SESSION_EXPIRE
		result.Message = "session decode err"
		return result
	}
	//update sso
	sso.Mtime = now.Unix()
	redis.HSet(key, "mtime", sso.Mtime)

	result.PushData("sso", sso)
	return result
}

func Remove(sessionId, token string, uid int64) *common.ResultPo {
	result := common.CreateResultPo()
	redis := database.Redis("sso")
	redis.Del(database.GetSSoSessionKey(sessionId))
	redis.HDel(database.SESSION_POOL, string(uid))
	return result
}

func GetSSo(sessionId, token string) (*SSoPo, error) {
	result := Validate(sessionId, token)
	if result.Code != 200 {
		return nil, fmt.Errorf("not found sso")
	}
	data := result.Data["sso"]
	sso := data.(SSoPo)
	return &sso, nil
}
