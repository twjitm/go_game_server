package user

import "go_game_server/grpcserver/server"

type IUserHandler interface {
	getUserById(uid int64)
}
type UserHandler struct {
	handler server.TcpHandler
}

func (handler UserHandler) Handler(data []byte) {

}

func (handler *UserHandler) getUserById(uid int64) {

}
