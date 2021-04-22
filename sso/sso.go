package sso

import "go_game_server/web"

type SSoPo struct {
	SessionId string
	Uid       int64
	Token     string
	Mtime     int64
}

func Register(po SSoPo) web.ResultPo {

}
